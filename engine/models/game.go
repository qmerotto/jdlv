package models

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type gamePool map[uuid.UUID]Game

var games = make(gamePool)

type Game struct {
	UUID     uuid.UUID `json:"uuid"`
	Grid     Grid      `json:"grid"`
	Rules    []Rule
	Running  bool
	UserUUID uuid.UUID `json:"userUUID"`
	cancel   context.CancelFunc
}

func NewGame(userUUID uuid.UUID, x, y int) (*Game, error) {
	gameUUID := uuid.New()
	ctx, cancel := context.WithCancel(context.Background())
	newGame := Game{
		UUID: gameUUID,
		Grid: NewGrid(x, y),
		Rules: []Rule{
			defaultRule,
		},
		Running:  false,
		UserUUID: userUUID,
		cancel:   cancel,
	}

	newGame.Run(ctx)
	games[gameUUID] = newGame

	return &newGame, nil
}

func StopGame(uuid uuid.UUID) {
	games[uuid].cancel()
}

func (g *Game) Run(ctx context.Context) {
	fmt.Println("game at work")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop engine")
			return
		case t := <-ticker.C:
			fmt.Println("Before actualize", t.UnixMicro())
			g.Grid.Actualize(g.Rules)
			fmt.Println("After actualize", time.Now().UnixMicro())
		default:
			time.Sleep(time.Second)
		}
	}
}
