package main

import (
	"context"
	"jdlv/engine"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go engine.Start(ctx)
	web.Run()
}
