package engine

import (
	"context"
	"testing"
)

func TestInit(t *testing.T) {
	INIT_GRID.toString()
}

func TestActualize(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	Start(ctx)
}
