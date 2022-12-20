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
	server.Run(ctx)
}
