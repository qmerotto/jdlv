package models

import (
	"fmt"
)

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

var currentGrid *Grid

var InitX = 100
var InitY = 100

func init() {
	grid := NewGrid()
	grid.build()
}

func NewGrid() *Grid {
	currentGrid = &Grid{
		X: InitX,
		Y: InitY,
	}
	return currentGrid
}

func CurrentGrid() *Grid {
	return currentGrid
}

func (g *Grid) build() {
	g.Raws = make([]Raw, InitY)

	for y := 0; y < InitY; y++ {
		g.Raws[y].Columns = make([]Column, InitX)
		for x := 0; x < InitX; x++ {
			g.Raws[y].Columns[x].Cell = DefaultCell(x, y, g)
		}
	}
}

func (g *Grid) Actualize() {
	tmpGrid := g.copy()

	for x := 0; x < InitX; x++ {
		for y := 0; y < InitY; y++ {
			g.Raw(y).Column(x).Actualize(tmpGrid.Raw(y).Column(x).AliveNeighbours())
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
	fmt.Printf(g.String())
}

func (g *Grid) String() string {
	s := ""

	for x := 0; x < InitX; x++ {
		for y := 0; y < InitY; y++ {
			if currentGrid.Raw(x).Column(y).State.Alive {
				s = fmt.Sprintf("%s%s", s, "X")
			} else {
				s = fmt.Sprintf("%s%s", s, " ")
			}
		}
		s = fmt.Sprintf("%s%s", s, "\n")
	}

	return s
}

func (g *Grid) copy() *Grid {
	tmpGrid := &Grid{}
	tmpGrid.Raws = make([]Raw, InitY)

	for y := 0; y < InitY; y++ {
		tmpGrid.Raws[y].Columns = make([]Column, InitX)
		for x := 0; x < InitX; x++ {
			tmpGrid.Raws[y].Columns[x].Cell = &Cell{
				State:  g.Raws[y].Columns[x].Cell.State,
				grid:   tmpGrid,
				row:    y,
				column: x,
			}
		}
	}

	for y := 0; y < InitY; y++ {
		for x := 0; x < InitX; x++ {
			tmpGrid.Raws[y].Columns[x].Cell.initNeighbours()
		}
	}

	return tmpGrid
}
