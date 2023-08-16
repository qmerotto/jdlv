package controllers

import (
	"encoding/json"
	"io/ioutil"
	"jdlv/engine"
	"jdlv/games/jdlv/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetGridOuput struct {
	Grid models.Grid `json:"grid"`
}

func GetGrid(c *gin.Context) {
	//c.JSON(200, GetGridOuput{Grid: models.CurrentGrid()})
}

type SetCellInput struct {
	GameUUID uuid.UUID `json:"gameUuid"`
	X        int       `json:"x"`
	Y        int       `json:"y"`
}

type SetCellsInput struct {
	Cells []SetCellInput `json:"cells"`
}

func SetCell(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	var setCellsInput SetCellInput

	if err = json.Unmarshal(body, &setCellsInput); err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	updateCell, err := engine.Instance().SetGridCell(engine.SetGridCellInput{
		GameUUID: setCellsInput.GameUUID,
		X:        setCellsInput.X,
		Y:        setCellsInput.Y,
	})

	if err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	c.JSON(200, updateCell)
}
