package main

import (
	"log"
	"os"

	"github.com/mmirolim/ws-fun/conf"
	ds "github.com/mmirolim/ws-fun/datastore"
	"github.com/mmirolim/ws-fun/ws"
)

func main() {
	// read config
	f, err := os.Open("conf.toml")
	fatalOnError(err)
	// read conf file
	AppConf, err := conf.Read(f)
	fatalOnError(err)
	// init datastore
	ds.Initialize(AppConf.DS.Redis.Port)
	msg := make(chan []byte)
	// start subscription listening
	go ds.Subscribe(AppConf.DS.Redis.Chan, msg)
	// start ws server
	err = ws.StartServer(AppConf.WS.Port)
	fatalOnError(err)
	for {
		select {
		case v := <-msg:
			// read from channel
			log.Printf("%v", string(v))
		}
	}

}

func fatalOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
