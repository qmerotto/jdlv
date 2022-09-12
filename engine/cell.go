package engine

import (
	"math/rand"
)

type Cell struct {
	grid       *Grid
	neighbours []*Cell
	row        int
	column     int
	State      CellState
}

type CellState struct {
	Alive       bool
	Fuel        int
	Temperature int
}

func (c *Cell) initState(grid *Grid) {
	c.State = CellState{
		Alive:       rand.Intn(2)%2 == 0,
		Fuel:        rand.Intn(100),
		Temperature: 20,
	}
	c.grid = grid

	c.initNeighbours()
}

func (c *Cell) initNeighbours() {
	//TODO 8 voisins...
	c.neighbours = make([]*Cell, 8)

	//N
	if c.row-1 >= 0 {
		c.neighbours[0] = c.grid.Raw(c.row - 1).Column(c.column)
	}

	//NE
	if c.row-1 >= 0 && c.column+1 < INIT_X {
		c.neighbours[1] = c.grid.Raw(c.row - 1).Column(c.column + 1)
	}

	//E
	if c.column+1 < INIT_X {
		c.neighbours[2] = c.grid.Raw(c.row).Column(c.column + 1)
	}

	//SE
	if c.row+1 < INIT_Y && c.column+1 < INIT_X {
		c.neighbours[3] = c.grid.Raw(c.row + 1).Column(c.column + 1)
	}

	//S
	if c.row+1 < INIT_Y {
		c.neighbours[4] = c.grid.Raw(c.row + 1).Column(c.column)
	}

	//SO
	if c.row+1 < INIT_Y && c.column-1 >= 0 {
		c.neighbours[5] = c.grid.Raw(c.row + 1).Column(c.column - 1)
	}

	//O
	if c.column-1 >= 0 {
		c.neighbours[6] = c.grid.Raw(c.row).Column(c.column - 1)
	}

	//NO
	if c.row-1 >= 0 && c.column-1 >= 0 {
		c.neighbours[7] = c.grid.Raw(c.row - 1).Column(c.column - 1)
	}
}

func (c *Cell) AliveNeighbours() int32 {
	aliveCount := 0
	neighbours := c.Neighbours()

	for i := 0; i < len(neighbours); i++ {
		neighbour := neighbours[i]
		if neighbour != nil && neighbour.Alive() {
			aliveCount++
		}
	}

	return int32(aliveCount)
}

func (c *Cell) Neighbours() []*Cell {
	return c.neighbours
}

func (c *Cell) Alive() bool {
	return c.State.Alive
}
