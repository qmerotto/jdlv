package server

import (
	"context"
	"fmt"
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

type WSMessage struct {
	Event      string      `json:"event"`
	Payload    interface{} `json:"payload"`
	UpdateTime time.Time   `json:"updateTime"`
}

func grid(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	ctx, cancel := context.WithCancel(r.Context())
	chWrite := make(chan []byte)

	go write(ctx, cancel, c, chWrite, time.Second)

	//chRead := make(chan []byte)
	//go read(ctx, cancel, c, chRead, time.Second)

	<-ctx.Done()
}

func read(ctx context.Context, cancel context.CancelFunc, c *websocket.Conn, bodyChan chan []byte, duration time.Duration) {
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		case msgBytes := <-bodyChan:

			if err := c.WriteMessage(2, msgBytes); err != nil {
				log.Println("write:", err)
				break
			}
		default:
			msgType, bytes, err := c.ReadMessage()
			if err != nil {
				fmt.Printf("reading error: %s, message type: %d", err, msgType)
			}

			bodyChan <- bytes
			time.Sleep(time.Second)
		}
	}

}

/*
tick := time.Tick(duration)

message := WSMessage{
	Event:      "gridUpdated",
	Payload:    "aaaaaaa",
	UpdateTime: t,
}

msgBytes, err := json.Marshal(message)
if err != nil {
	log.Println("write:", err)
	break
}*/

func write(ctx context.Context, cancel context.CancelFunc, c *websocket.Conn, bodyChan chan []byte, duration time.Duration) {
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		case msgBytes := <-bodyChan:
			if err := c.WriteMessage(2, msgBytes); err != nil {
				log.Println("write:", err)
				break
			}
		default:
			time.Sleep(time.Second)
		}
	}
}
