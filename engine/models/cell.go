package models

type Cell struct {
	grid       Grid
	neighbours []*Cell
	X          int       `json:"x"`
	Y          int       `json:"y"`
	State      CellState `json:"state"`
}

type CellState struct {
	Alive       bool `json:"alive"`
	Fuel        uint `json:"fuel"`
	Temperature int  `json:"temperature"`
}

func LivingCell(x, y int, g Grid) Cell {
	return Cell{
		X:    x,
		Y:    y,
		grid: g,
		State: CellState{
			Alive:       true,
			Fuel:        0,
			Temperature: 20,
		},
	}
}

func DefaultCell(x, y int, g Grid) Cell {
	return Cell{
		X:    x,
		Y:    y,
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
	if c.Y-1 >= 0 {
		c.neighbours[0] = &c.grid[c.X][c.Y-1]
	}

	//NE
	if c.Y-1 >= 0 && c.X+1 < InitX {
		c.neighbours[1] = &c.grid[c.X+1][c.Y-1]
	}

	//E
	if c.X+1 < InitX {
		c.neighbours[2] = &c.grid[c.X+1][c.Y]
	}

	//SE
	if c.Y+1 < InitY && c.X+1 < InitX {
		c.neighbours[3] = &c.grid[c.X+1][c.Y+1]
	}

	//S
	if c.Y+1 < InitY {
		c.neighbours[4] = &c.grid[c.X][c.Y+1]
	}

	//SO
	if c.Y+1 < InitY && c.X-1 >= 0 {
		c.neighbours[5] = &c.grid[c.X-1][c.Y+1]
	}

	//O
	if c.X-1 >= 0 {
		c.neighbours[6] = &c.grid[c.X-1][c.Y]
	}

	//NO
	if c.Y-1 >= 0 && c.X-1 >= 0 {
		c.neighbours[7] = &c.grid[c.X-1][c.Y-1]
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

func (c *Cell) Actualize() Cell {
	res := Cell{
		grid:       c.grid,
		neighbours: c.neighbours,
		X:          c.X,
		Y:          c.Y,
		State:      c.State,
	}

	n := c.AliveNeighbours()
	if n < 1 {
		res.State.Alive = false
	}

	if !c.Alive() && n >= 3 {
		res.State.Alive = true
	}

	if n >= 5 {
		res.State.Alive = false
	}

	return res
}
