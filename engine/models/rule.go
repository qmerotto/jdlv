package models

type Rule func(c *Cell, args ...interface{})

func defaultRule(c *Cell, _ ...interface{}) {
	n := c.AliveNeighbours()
	if c.Alive() && (n == 2 || n == 3) {
		c.State.Alive = true
	} else if !c.Alive() && n == 3 {
		c.State.Alive = true
	} else {
		c.State.Alive = false
	}
}
