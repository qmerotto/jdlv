package engine

import (
	"context"
	"fmt"
	"jdlv/games/jdlv"
	"time"

	"github.com/google/uuid"
)

var instance engine

type engine struct {
	running bool
	start   chan struct{}
	stop    chan struct{}
}

type Runnable[T, V any] interface {
	Uuid() uuid.UUID
	Start(ctx context.Context) error
}

type gamePool map[uuid.UUID]Runnable[interface{}, interface{}]

var games = make(gamePool)
var gameChan = make(chan Runnable[interface{}, interface{}])
var gameCancels = make(map[uuid.UUID]context.CancelFunc, 0)

func init() {
	instance = engine{
		running: false,
		start:   make(chan struct{}),
		stop:    make(chan struct{}),
	}
}

func Instance() *engine {
	return &instance
}

func (e *engine) Start() error {
	if e.running {
		return fmt.Errorf("already started")
	}

	e.start <- struct{}{}
	e.running = true

	return nil
}

func (e *engine) Stop() error {
	if !e.running {
		return fmt.Errorf("not running")
	}

	e.stop <- struct{}{}
	e.running = false

	return nil
}

func (e *engine) Run(ctx context.Context) {
	for {
		<-e.start
		currentCtx, cancel := context.WithCancel(ctx)
		go e.work(currentCtx)
		<-e.stop
		cancel()
	}
}

func (e *engine) work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case game := <-gameChan:
			gameCtx, gameCancel := context.WithCancel(ctx)
			games[game.Uuid()] = game
			gameCancels[game.Uuid()] = gameCancel
			go game.Start(gameCtx)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (e *engine) cancelAllGames() {
	for _, cancel := range gameCancels {
		cancel()
	}
}

func (e *engine) IsRunning() bool {
	return e.running
}

type NewGameInput struct {
	UserUUID uuid.UUID `json:"userUUID"`
	X        int       `json:"x"`
	Y        int       `json:"y"`
}

func (e *engine) NewGame(params NewGameInput) (interface{}, error) {
	jdlv, err := jdlv.New(context.Background(), jdlv.InputNew{
		UserUUID: params.UserUUID,
		X:        params.X,
		Y:        params.Y,
	})
	if err != nil {
		return nil, err
	}

	games[jdlv.UUID] = jdlv

	return jdlv, nil
}

func (e *engine) StartGame(uuid uuid.UUID) error {
	if games[uuid] != nil {
		gameChan <- games[uuid]
		return nil
	}

	return fmt.Errorf("games %s doesnt exist", uuid)
}

func (e *engine) StopGame(uuid uuid.UUID) error {
	if gameCancels[uuid] != nil {
		gameCancels[uuid]()
		return nil
	}

	return fmt.Errorf("games %s doesnt exist", uuid)
}

type setCellInput struct {
	GameUUID uuid.UUID `json:"gameUUID"`
	UserUUID uuid.UUID `json:"userUUID"`
	X        int       `json:"x"`
	Y        int       `json:"y"`
}

func (e *engine) SetGrid(params setCellInput) error {
	if game := games[params.GameUUID]; game != nil {
		if jdlvGame, ok := game.(*jdlv.Game); ok {
			jdlvGame.SetCell(jdlv.SetCellInput{
				UserUUID: params.UserUUID,
				X:        params.X,
				Y:        params.Y,
			})
		}
	}

	return nil
}
