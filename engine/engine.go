package engine

import (
	"context"
	"fmt"
	"jdlv/games/jdlv"
	"jdlv/games/jdlv/models"
	"time"

	"github.com/google/uuid"
)

var instance engine

type engine struct {
	running bool
	start   chan struct{}
	stop    chan struct{}
}

type Runnable interface {
	Uuid() uuid.UUID
	Start(ctx context.Context, output chan []byte) error
}

type gamePool[T any] map[uuid.UUID]T

type gameChanMsg struct {
	runnable Runnable
	output   chan []byte
}

var tokensMap = make(map[uuid.UUID]uuid.UUID)
var games = make(gamePool[Runnable])
var gameChan = make(chan gameChanMsg)
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
	tick := time.Tick(5 * time.Second)
	fmt.Printf("engine starting to work...\n")
	for {
		select {
		case <-ctx.Done():
			return
		case game := <-gameChan:
			gameCtx, gameCancel := context.WithCancel(ctx)
			//games[game.runnable.Uuid()] = game.runnable
			gameCancels[game.runnable.Uuid()] = gameCancel

			fmt.Printf("starting game %s\n", game.runnable.Uuid())
			go game.runnable.Start(gameCtx, game.output)
		case <-tick:
			fmt.Printf("Current Games: %v\n", games)
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

func (e *engine) GetGameByToken(token string) {

}

type NewGameInput struct {
	UserUUID uuid.UUID `json:"userUUID"`
	X        int       `json:"x"`
	Y        int       `json:"y"`
}

func (e *engine) NewGame(params NewGameInput) (Runnable, error) {
	jdlv, err := jdlv.New(context.Background(), jdlv.InputNew{
		UserUUID: params.UserUUID,
		X:        params.X,
		Y:        params.Y,
	})
	if err != nil {
		return nil, err
	}

	fmt.Printf("Game %s successfully created !\n", jdlv.UUID)
	games[jdlv.UUID] = jdlv

	return jdlv, nil
}

func (e *engine) CreateGameToken(gameUUID uuid.UUID) (*uuid.UUID, error) {
	if games[gameUUID] != nil {
		token := uuid.New()

		fmt.Printf("New Token: %s\n", token)
		// TODO Store hashed token in DB
		tokensMap[token] = gameUUID
		return &token, nil
	}

	return nil, fmt.Errorf("games %s doesnt exist", gameUUID.String())
}

func (e *engine) StartGame(token uuid.UUID, cancel context.CancelFunc, output chan []byte) error {
	gameUUID := tokensMap[token]
	if games[gameUUID] != nil {
		gameChan <- gameChanMsg{
			runnable: games[gameUUID],
			output:   output,
		}

		fmt.Printf("games %s started !", gameUUID.String())
		return nil
	}

	gameCancels[gameUUID] = cancel

	return fmt.Errorf("games %s doesnt exist", gameUUID.String())
}

func (e *engine) StopGame(uuid uuid.UUID) error {
	if gameCancels[uuid] != nil {
		gameCancels[uuid]()
		return nil
	}

	return fmt.Errorf("game %s doesnt exist", uuid)
}

type SetGridCellInput struct {
	GameUUID uuid.UUID `json:"gameUuid"`
	UserUUID uuid.UUID `json:"userUuid"`
	X        int       `json:"x"`
	Y        int       `json:"y"`
}

func (e *engine) SetGridCell(params SetGridCellInput) (*models.Cell, error) {
	if game := games[params.GameUUID]; game != nil {
		if jdlvGame, ok := game.(*jdlv.Game); ok {
			updatedCell := jdlvGame.SetCell(jdlv.SetCellInput{
				UserUUID: params.UserUUID,
				X:        params.X,
				Y:        params.Y,
			})
			//fmt.Printf("updated grid %v\n", jdlvGame.Grid[params.X][params.Y])
			return &updatedCell, nil
		}
		return nil, fmt.Errorf("invalid game type")
	}

	return nil, fmt.Errorf("game %s doesnt exist", params.GameUUID)
}
