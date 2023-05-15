package controllers

import (
	"encoding/json"
	"io/ioutil"
	"jdlv/engine"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	newGame, err := engine.Instance().NewGame(engine.NewGameInput{
		UserUUID: userUUID,
		X:        newGameParameters.X,
		Y:        newGameParameters.Y,
	})
	if err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	c.JSON(200, gin.H{"newGame": newGame})
}

type startInput struct {
	GameUUID uuid.UUID `json:"gameUUID"`
}

func StartGame(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	var newStartParameters startInput
	if err = json.Unmarshal(body, &newStartParameters); err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	err = engine.Instance().StartGame(newStartParameters.GameUUID)
	if err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	return
}

type stopInput struct {
	gameUUID uuid.UUID
}

func StopGame(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	var newStopParameters stopInput
	if err = json.Unmarshal(body, &newStopParameters); err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	err = engine.Instance().StopGame(newStopParameters.gameUUID)
	if err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}
}

func JDLVSetCell(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

	var newStopParameters stopInput
	if err = json.Unmarshal(body, &newStopParameters); err != nil {
		log.Printf(err.Error())
		c.AbortWithStatus(500)
	}

}
