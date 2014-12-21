package ws

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
)

func TestStart(t *testing.T) {
	fmt.Println("START TESTING")
}

func TestNewWsServer(t *testing.T) {
	// start ws server
	fmt.Println("Starting Server ws")
	err := NewWsServer(8081, "/echo")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	origin := "http://localhost/"
	url := "http://localhost:12345/echo"
	dialer := websocket.Dialer{}
	var header http.Header
	header["Origin"] = []string{origin}
	c, res, err := dialer.Dial(url, header)
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", c)
}
