package engine

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"jdlv/engine/models"
	"time"
)

var instance engine

type engine struct {
	games   chan *models.Game
	cancels map[uuid.UUID]context.CancelFunc
	running bool
	start   chan struct{}
	stop    chan struct{}
}

func init() {
	instance = engine{
		games:   make(chan *models.Game),
		cancels: map[uuid.UUID]context.CancelFunc{},
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

func (e *engine) StartGame(game *models.Game) error {
	e.games <- game
	game.Running = true

	return nil
}

func (e *engine) StopGame(game *models.Game) error {
	cancelFunc, ok := e.cancels[game.UUID]
	if ok {
		cancelFunc()
		delete(e.cancels, game.UUID)
		return nil
	}

	return fmt.Errorf("stop game error")
}

func (e *engine) work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case game := <-e.games:
			gameCtx, gameCancel := context.WithCancel(ctx)
			e.cancels[game.UUID] = gameCancel
			go game.Run(gameCtx)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (e *engine) cancelAllGames() {
	for _, cancel := range e.cancels {
		cancel()
	}
}
func (e *engine) IsRunning() bool {
	return e.running
}
