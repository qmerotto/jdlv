package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

type WSMEssage struct {
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
}

func grid(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for i := 0; true; i++ {

		message := WSMEssage{
			Event:   "gridUpadted",
			Payload: "aaaaaaa",
		}
		msgBytes, err := json.Marshal(message)
		if err != nil {
			log.Println("write:", err)
			break
		}

		err = c.WriteMessage(2, msgBytes)
		if err != nil {
			log.Println("write:", err)
			break
		}
		time.Sleep(time.Second)
	}
}
