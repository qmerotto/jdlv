package engine

import (
	"context"
	"jdlv/engine/models"
	"testing"
)

func TestInit(t *testing.T) {
	models.CurrentGrid().String()
}

func TestActualize(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	Start(ctx)
}
