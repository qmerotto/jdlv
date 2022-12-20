package models

import (
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AliveCellsSuite struct {
	suite.Suite
	LivingCell Cell
}

func (suite *AliveCellsSuite) SetupSuite() {}

func (suite *AliveCellsSuite) SetupTest() {
	suite.LivingCell = Cell{State: CellState{Alive: true}}
}

func TestAliveCellsSuite(t *testing.T) {
	suite.Run(t, new(AliveCellsSuite))
}

func (suite *AliveCellsSuite) TestIsolation() {
	suite.LivingCell.neighbours = livingCells(1)
	defaultRule(&suite.LivingCell)
	assert.Equal(suite.T(), false, suite.LivingCell.State.Alive)
}

func (suite *AliveCellsSuite) TestOverpopulation() {
	suite.LivingCell.neighbours = livingCells(4)
	defaultRule(&suite.LivingCell)
	assert.Equal(suite.T(), false, suite.LivingCell.State.Alive)

	suite.LivingCell.neighbours = livingCells(5)
	defaultRule(&suite.LivingCell)
	assert.Equal(suite.T(), false, suite.LivingCell.State.Alive)

	suite.LivingCell.neighbours = livingCells(6)
	defaultRule(&suite.LivingCell)
	assert.Equal(suite.T(), false, suite.LivingCell.State.Alive)

	suite.LivingCell.neighbours = livingCells(7)
	defaultRule(&suite.LivingCell)
	assert.Equal(suite.T(), false, suite.LivingCell.State.Alive)

	suite.LivingCell.neighbours = livingCells(8)
	defaultRule(&suite.LivingCell)
	assert.Equal(suite.T(), false, suite.LivingCell.State.Alive)
}

type DeadCellsSuite struct {
	suite.Suite
	DeadCell Cell
}

func (suite *DeadCellsSuite) SetupSuite() {}

func (suite *DeadCellsSuite) SetupTest() {
	suite.DeadCell = Cell{State: CellState{Alive: false}}
}

func TestDeadCellsSuite(t *testing.T) {
	suite.Run(t, new(DeadCellsSuite))
}

func (suite *DeadCellsSuite) TestRevival() {
	suite.DeadCell.neighbours = livingCells(3)
	defaultRule(&suite.DeadCell)
	assert.Equal(suite.T(), true, suite.DeadCell.State.Alive)
}

func (suite *DeadCellsSuite) TestUnchanged() {
	suite.DeadCell.neighbours = livingCells(1)
	defaultRule(&suite.DeadCell)
	assert.Equal(suite.T(), false, suite.DeadCell.State.Alive)

	suite.DeadCell.neighbours = livingCells(2)
	defaultRule(&suite.DeadCell)
	assert.Equal(suite.T(), false, suite.DeadCell.State.Alive)

	suite.DeadCell.neighbours = livingCells(4)
	defaultRule(&suite.DeadCell)
	assert.Equal(suite.T(), false, suite.DeadCell.State.Alive)

	suite.DeadCell.neighbours = livingCells(5)
	defaultRule(&suite.DeadCell)
	assert.Equal(suite.T(), false, suite.DeadCell.State.Alive)

	suite.DeadCell.neighbours = livingCells(6)
	defaultRule(&suite.DeadCell)
	assert.Equal(suite.T(), false, suite.DeadCell.State.Alive)

	suite.DeadCell.neighbours = livingCells(7)
	defaultRule(&suite.DeadCell)
	assert.Equal(suite.T(), false, suite.DeadCell.State.Alive)

	suite.DeadCell.neighbours = livingCells(8)
	defaultRule(&suite.DeadCell)
	assert.Equal(suite.T(), false, suite.DeadCell.State.Alive)
}

func livingCells(n int) []*Cell {
	cells := []*Cell{
		{State: CellState{Alive: false}},
		{State: CellState{Alive: false}},
		{State: CellState{Alive: false}},
		{State: CellState{Alive: false}},
		{State: CellState{Alive: false}},
		{State: CellState{Alive: false}},
		{State: CellState{Alive: false}},
		{State: CellState{Alive: false}},
	}

	for i := 0; i < n; i++ {
		cells[i].State.Alive = true
	}

	return cells
}
