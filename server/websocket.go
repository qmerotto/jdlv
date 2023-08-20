package server

import (
	"context"
	"encoding/json"
	"fmt"
	"jdlv/engine"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
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

	ctx := r.Context()
	token, err := readAuth(ctx, c)
	if err != nil {
		panic(err)
	}
	if token == nil {
		return
	}

	if token != nil {
		ctx, cancel := context.WithCancel(ctx)

		chWrite := make(chan []byte)
		go write(ctx, c, chWrite, time.Second)

		if err = engine.Instance().StartGame(*token, cancel, chWrite); err != nil {
			return
		}
	}

	<-ctx.Done()
}

type authMessage struct {
	Token string `json:"token"`
}

func readAuth(ctx context.Context, c *websocket.Conn) (*uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, nil
		default:
			msgType, bytes, err := c.ReadMessage()
			if err != nil {
				fmt.Printf("reading error: %s, message type: %d", err, msgType)
				return nil, err
			}

			if len(bytes) > 0 {
				var authMsg authMessage
				if err := json.Unmarshal(bytes, &authMsg); err != nil {
					return nil, err
				}

				fmt.Printf("auth message: %v\n", authMsg)
				token := uuid.MustParse(authMsg.Token)
				return &token, nil
			}

			time.Sleep(100 * time.Millisecond)
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

func write(ctx context.Context, c *websocket.Conn, bodyChan chan []byte, period time.Duration) {

	for {
		select {
		case <-ctx.Done():
			return
		case msgBytes := <-bodyChan:
			fmt.Printf("writing to websocket")
			if err := c.WriteMessage(2, msgBytes); err != nil {
				log.Println("write:", err)
				break
			}
		default:
			time.Sleep(period)
		}
	}
}
