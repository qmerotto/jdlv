package engine

import (
	"context"
	"fmt"
	"jdlv/engine/models"
	"time"
)

var start = make(chan struct{})
var stop = make(chan struct{})
var running = false

func Start() error {
	if running {
		return fmt.Errorf("already started")
	}

	start <- struct{}{}
	running = true

	return nil
}

func Stop() error {
	if !running {
		return fmt.Errorf("not running")
	}

	stop <- struct{}{}
	running = false

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

func IsRunning() bool {
	return running
}

func work(ctx context.Context) {
	fmt.Println("engine at work")

	ticker := time.NewTicker(5 * time.Second)

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
