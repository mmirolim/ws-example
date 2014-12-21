package ws

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func srvWs(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		msgType, p, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(msgType, p)
		if err = c.WriteMessage(websocket.TextMessage, []byte("REPLY")); err != nil {
			log.Println(err)
			return
		}
	}
}

func NewWsServer(port int, route string) error {
	p := ":" + strconv.Itoa(port)
	log.Println("starting ws handler")
	http.HandleFunc(route, srvWs)
	err := http.ListenAndServe(p, nil)
	if err != nil {
		log.Println(err)
	}
	return err
}
