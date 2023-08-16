package jdlv

import (
	"context"
	"encoding/json"
	"fmt"
	"jdlv/games/jdlv/models"
	"time"

	"github.com/google/uuid"
)

type Game struct {
	UUID     uuid.UUID     `json:"uuid"`
	Grid     models.Grid   `json:"grid"`
	Rules    []models.Rule `json:"-"`
	Running  bool          `json:"-"`
	UserUUID uuid.UUID     `json:"userUuid"`
}

type InputNew struct {
	UserUUID uuid.UUID `json:"userUUID"`
	X        int       `json:"x"`
	Y        int       `json:"y"`
}

func New(ctx context.Context, opts ...interface{}) (*Game, error) {
	var input InputNew

	if len(opts) != 1 {
		return nil, fmt.Errorf("invalid opt length: %d", len(opts))
	}
	if _, ok := opts[0].(InputNew); ok {
		input = opts[0].(InputNew)
	} else {
		return nil, fmt.Errorf("invalid input")
	}

	newgame := Game{
		UUID: uuid.New(),
		Grid: models.NewGrid(input.X, input.Y),
		Rules: []models.Rule{
			models.DefaultRule,
		},
		Running:  false,
		UserUUID: input.UserUUID,
	}

	return &newgame, nil
}

func (g *Game) Uuid() uuid.UUID {
	return g.UUID
}

func (g *Game) Start(ctx context.Context, output chan []byte) error {
	defer func() error {
		if e := recover(); e != nil {
			return fmt.Errorf("%v", e)
		}

		return nil
	}()

	g.run(ctx, output)

	return nil
}

func (g *Game) run(ctx context.Context, output chan []byte) {
	fmt.Println("game at work")
	g.Running = true

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop game")
			return
		case t := <-ticker.C:
			fmt.Println("Before actualize", t.UnixMicro())
			g.Grid.Actualize(g.Rules)
			output <- g.dispatchGridUpdatedBody()
			fmt.Println("After actualize", time.Now().UnixMicro())
		default:
			time.Sleep(time.Second)
		}
	}
}

type GridErrorEvent struct {
	Event string `json:"event"`
}

type GridUpdatedEvent struct {
	Event string      `json:"event"`
	Grid  models.Grid `json:"grid"`
}

func (g *Game) dispatchGridUpdatedBody() []byte {
	event := GridUpdatedEvent{
		Event: "gridUpdated",
		Grid:  g.Grid,
	}
	bodyBytes, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}

	return bodyBytes
}

type SetCellInput struct {
	UserUUID uuid.UUID `json:"userUUID"`
	X        int       `json:"x"`
	Y        int       `json:"y"`
}

func (g *Game) SetCell(params SetCellInput) models.Cell {
	g.Grid[params.X][params.Y].State.Alive = true
	return g.Grid[params.X][params.Y]
}
