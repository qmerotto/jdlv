package engine

import (
	"context"
	"fmt"
	"jdlv/engine/models"
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
				models.CurrentGrid().Actualize()
				fmt.Println("Tick at", t)
			}
		}
	}()

	wg.Wait()
}
