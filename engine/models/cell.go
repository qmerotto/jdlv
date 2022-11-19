package models

type Cell struct {
	grid       Grid
	neighbours []*Cell
	y          int
	x          int
	State      CellState
}

type CellState struct {
	Alive       bool
	Fuel        int
	Temperature int
}

func DefaultCell(x, y int, g Grid) Cell {
	return Cell{
		y:    y,
		x:    x,
		grid: g,
		State: CellState{
			Alive:       false,
			Fuel:        0,
			Temperature: 20,
		},
	}
}

func (c *Cell) initNeighbours() {
	//TODO 8 voisins...
	c.neighbours = make([]*Cell, 8)

	//N
	if c.y-1 >= 0 {
		c.neighbours[0] = &c.grid[c.x][c.y-1]
	}

	//NE
	if c.y-1 >= 0 && c.x+1 < InitX {
		c.neighbours[1] = &c.grid[c.x+1][c.y-1]
	}

	//E
	if c.x+1 < InitX {
		c.neighbours[2] = &c.grid[c.x+1][c.y]
	}

	//SE
	if c.y+1 < InitY && c.x+1 < InitX {
		c.neighbours[3] = &c.grid[c.x+1][c.y+1]
	}

	//S
	if c.y+1 < InitY {
		c.neighbours[4] = &c.grid[c.x][c.y+1]
	}

	//SO
	if c.y+1 < InitY && c.x-1 >= 0 {
		c.neighbours[5] = &c.grid[c.x-1][c.y+1]
	}

	//O
	if c.x-1 >= 0 {
		c.neighbours[6] = &c.grid[c.x-1][c.y]
	}

	//NO
	if c.y-1 >= 0 && c.x-1 >= 0 {
		c.neighbours[7] = &c.grid[c.x-1][c.y-1]
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

func (c *Cell) Actualize(n int32) {

	if n < 1 {
		c.State.Alive = false
	}

	if !c.Alive() && n >= 3 {
		c.State.Alive = true
	}

	if n >= 5 {
		c.State.Alive = false
	}
}
