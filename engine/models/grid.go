package models

import (
	"fmt"
)

type Grid [][]Cell

var InitX = 10
var InitY = 10

func NewGrid(x, y int) Grid {
	newGrid := make(Grid, x)
	for x := 0; x < InitX; x++ {
		newGrid[x] = make([]Cell, y)
	}

	return newGrid
}

func (g Grid) build() {
	for x := 0; x < InitX; x++ {
		for y := 0; y < InitY; y++ {
			g[x][y] = DefaultCell(x, y, g)
		}
	}

	g.initCellsNeighbours()

	fmt.Print("build\n")
	fmt.Print(g.String())
}

func (g Grid) Reinitialize() {
	g.build()
}

func (g Grid) Actualize(rules []Rule) {
	copyGrid := g.copy()

	for x := 0; x < InitX; x++ {
		for y := 0; y < InitY; y++ {
			g[x][y] = copyGrid[x][y].Actualize(rules)
		}
	}
	g.Print()
}

func (g Grid) Print() {
	fmt.Printf(g.String())
}

func (g Grid) String() string {
	s := ""

	for x := 0; x < InitX; x++ {
		for y := 0; y < InitY; y++ {
			if g[x][y].State.Alive {
				s = fmt.Sprintf("%s%s", s, "X")
			} else {
				s = fmt.Sprintf("%s%s", s, " ")
			}
		}
		s = fmt.Sprintf("%s%s", s, "\n")
	}

	return s
}

func (g Grid) copy() Grid {
	copyGrid := make(Grid, InitX)

	for x := 0; x < InitX; x++ {
		copyGrid[x] = make([]Cell, InitY)
		for y := 0; y < InitY; y++ {
			copyGrid[x][y] = Cell{
				X:     g[x][y].X,
				Y:     g[x][y].Y,
				State: g[x][y].State,
				grid:  copyGrid,
			}
		}
	}

	copyGrid.initCellsNeighbours()

	return copyGrid
}

func (g Grid) initCellsNeighbours() {
	for y := 0; y < InitY; y++ {
		for x := 0; x < InitX; x++ {
			g[x][y].initNeighbours()
		}
	}
}
