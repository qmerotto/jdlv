package models

import (
	"fmt"
)

type Grid [][]Cell

var currentGrid = make(Grid, 0)
var InitX = 100
var InitY = 100

func init() {
	grid := NewGrid()
	grid.build()
}

func NewGrid() Grid {
	if len(currentGrid) > 0 {
		return currentGrid
	}

	currentGrid = make(Grid, InitX)
	for x := 0; x < InitX; x++ {
		currentGrid[x] = make([]Cell, InitY)
	}

	return currentGrid
}

func CurrentGrid() Grid {
	return currentGrid
}

func (g Grid) build() {
	for x := 0; x < InitX; x++ {
		for y := 0; y < InitY; y++ {
			g[x][y] = DefaultCell(x, y, g)
		}
	}

	for x := 0; x < InitX; x++ {
		for y := 0; y < InitY; y++ {
			g[x][y].initNeighbours()
		}
	}

	fmt.Print("build\n")
	fmt.Print(g.String())
}

func (g Grid) Actualize() {
	tmpGrid := g.copy()

	for x := 0; x < InitX; x++ {
		for y := 0; y < InitY; y++ {
			g[x][y].Actualize(tmpGrid[x][y].AliveNeighbours())
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
			if currentGrid[x][y].State.Alive {
				s = fmt.Sprintf("%s%s", s, "X")
			} else {
				s = fmt.Sprintf("%s%s", s, " ")
			}
		}
		s = fmt.Sprintf("%s%s", s, "\n")
	}

	return s
}

func (g *Grid) Serialize() {

}

func (g Grid) copy() Grid {
	tmpGrid := make(Grid, InitX)

	for x := 0; x < InitX; x++ {
		tmpGrid[x] = make([]Cell, InitY)
		for y := 0; y < InitY; y++ {
			tmpGrid[x][y] = Cell{
				State: g[x][y].State,
				grid:  tmpGrid,
			}
		}
	}

	for y := 0; y < InitY; y++ {
		for x := 0; x < InitX; x++ {
			tmpGrid[x][y].initNeighbours()
		}
	}

	return tmpGrid
}
