package main

import (
	"context"
	"jdlv/engine"
	"jdlv/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	instance := engine.Instance()
	go instance.Run(ctx)
	if err := instance.Start(); err != nil {
		panic(err)
	}

	server.Run(ctx)
}
