package main

import (
	"context"
	"jdlv/engine"
	"jdlv/server"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go engine.Run(ctx)
	server.Run(ctx)
}
