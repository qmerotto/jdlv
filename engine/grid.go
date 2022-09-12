package engine

import "fmt"

type Grid struct {
	X    int
	Y    int
	Raws []Raw
}

type Raw struct {
	Columns []Column
}

type Column struct {
	Cell *Cell
}

var INIT_GRID *Grid
var INIT_X int = 10
var INIT_Y int = 10

func init() {
	grid := NewGrid()
	grid.Raws = make([]Raw, INIT_Y)

	for y := 0; y < INIT_Y; y++ {
		grid.Raws[y].Columns = make([]Column, INIT_X)
		for x := 0; x < INIT_X; x++ {
			grid.Raws[y].Columns[x].Cell = &Cell{
				row:    y,
				column: x,
				grid:   grid}
		}
	}

	for y := 0; y < INIT_Y; y++ {
		for x := 0; x < INIT_X; x++ {
			grid.Raws[y].Columns[x].Cell.initState(grid)
		}
	}

	grid.Print()
}

func NewGrid() *Grid {
	INIT_GRID = &Grid{
		X: INIT_X,
		Y: INIT_Y,
	}
	return INIT_GRID
}

func (g *Grid) Actualize() {
	tmpGrid := g.Copy()

	for y := 0; y < INIT_Y; y++ {
		for x := 0; x < INIT_X; x++ {
			aliveNeighbours := tmpGrid.Raw(y).Column(x).AliveNeighbours()
			cell := g.Raw(y).Column(x)
			if aliveNeighbours < 1 {
				cell.State.Alive = false
			}

			if !cell.Alive() && aliveNeighbours >= 3 {
				cell.State.Alive = true
			}

			if aliveNeighbours >= 5 {
				cell.State.Alive = false
			}

		}
	}
	g.Print()
}

func (g *Grid) Raw(x int) *Raw {
	return &g.Raws[x]
}

func (r *Raw) Column(y int) *Cell {
	return r.Columns[y].Cell
}

func (g *Grid) Print() {
	fmt.Printf(g.toString())
}

func (g *Grid) toString() string {
	s := ""

	for x := 0; x < INIT_X; x++ {
		raw := INIT_GRID.Raw(x)
		for y := 0; y < INIT_Y; y++ {
			cell := raw.Column(y)
			if cell.State.Alive {
				s = fmt.Sprintf("%s%s", s, "X")
			} else {
				s = fmt.Sprintf("%s%s", s, " ")
			}
		}
		s = fmt.Sprintf("%s%s", s, "\n")
	}

	return s
}

func (g *Grid) Copy() *Grid {
	tmpGrid := &Grid{}
	tmpGrid.Raws = make([]Raw, INIT_Y)

	for y := 0; y < INIT_Y; y++ {
		tmpGrid.Raws[y].Columns = make([]Column, INIT_X)
		for x := 0; x < INIT_X; x++ {
			tmpGrid.Raws[y].Columns[x].Cell = &Cell{
				State:  g.Raws[y].Columns[x].Cell.State,
				grid:   tmpGrid,
				row:    y,
				column: x,
			}
		}
	}

	for y := 0; y < INIT_Y; y++ {
		for x := 0; x < INIT_X; x++ {
			tmpGrid.Raws[y].Columns[x].Cell.initNeighbours()
		}
	}

	return tmpGrid
}
