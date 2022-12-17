package controllers

import (
	"encoding/json"
	"io/ioutil"
	"jdlv/engine/models"
	"log"

	"github.com/gin-gonic/gin"
)

type GetGridOuput struct {
	Grid models.Grid `json:"grid"`
}

func GetGrid(c *gin.Context) {
	c.JSON(200, GetGridOuput{Grid: models.CurrentGrid()})
}

type cellsInput struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type SetCellsInput struct {
	Cells []cellsInput `json:"cells"`
}

func SetCells(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	var setCellsInput SetCellsInput

	if err = json.Unmarshal(body, &setCellsInput); err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	grid := models.CurrentGrid()
	for _, cell := range setCellsInput.Cells {
		grid[cell.X][cell.Y].State.Alive = true
	}

	c.JSON(200, []byte(`{"status": "ok"}`))
}
