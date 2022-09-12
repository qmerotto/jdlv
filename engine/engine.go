package engine

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Start(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case t := <-ticker.C:
				INIT_GRID.Actualize()
				fmt.Println("Tick at", t)
			}
		}
	}()

	wg.Wait()
}
