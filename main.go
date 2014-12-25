package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mmirolim/ws-fun/conf"
	ds "github.com/mmirolim/ws-fun/datastore"
	r "github.com/mmirolim/ws-fun/routes"
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
	go func() {
		for v := range msg {
			log.Println(v)
		}
	}()
	// get routes
	mux := r.Initialize()
	// start routes
	err = http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatalln(err)
	}

}

func fatalOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
