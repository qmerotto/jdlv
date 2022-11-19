package engine

import (
	"context"
	"fmt"
	"jdlv/engine/models"
	"time"
)

var start = make(chan struct{}, 1)
var stop = make(chan struct{}, 1)
var Running = false

func Start() error {
	if Running {
		return fmt.Errorf("cannot start")
	}

	start <- struct{}{}
	Running = true

	return nil
}

func Stop() error {
	if !Running {
		return fmt.Errorf("cannot stop")
	}

	stop <- struct{}{}
	Running = false

	return nil
}

func Run(ctx context.Context) {
	for {
		<-start
		currentCtx, cancel := context.WithCancel(ctx)
		go work(currentCtx)
		<-stop
		cancel()
	}
}

func work(ctx context.Context) {
	fmt.Println("engine at work")

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop engine")
			return
		case t := <-ticker.C:
			models.CurrentGrid().Actualize()
			fmt.Println("Tick at", t)
		default:
			time.Sleep(100 * time.Microsecond)
		}
	}
}
