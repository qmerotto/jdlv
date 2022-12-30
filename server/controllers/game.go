package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"jdlv/engine"
	"jdlv/engine/models"
	"log"
)

var userUUID = uuid.MustParse("eebf381a-5b39-4eb9-9794-06bcae6c766e")

type newGameInput struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewGame(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	var newGameParameters newGameInput
	if err = json.Unmarshal(body, &newGameParameters); err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	newGame, err := models.NewGame(userUUID, newGameParameters.X, newGameParameters.Y)
	if err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	err = engine.Instance().StartGame(newGame)
	if err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	response, err := json.Marshal(newGame)
	if err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	c.JSON(200, response)
}
