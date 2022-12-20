package models

import (
	"context"
	"fmt"
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type Game struct {
	UUID    uuid.UUID
	Grid    Grid
	Rules   []Rule
	Running bool
	User    uuid.UUID
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
			time.Sleep(100 * time.Microsecond)
		}
	}
}
