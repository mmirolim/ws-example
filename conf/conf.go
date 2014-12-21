package conf

import (
	"io"

	"github.com/BurntSushi/toml"
)

type Datastore struct {
	Redis struct {
		Port int
		Chan string
	}
}

type Websocket struct {
	Port int
}

type App struct {
	DS Datastore
	WS Websocket
}

func Read(r io.Reader) (App, error) {
	var conf App
	_, err := toml.DecodeReader(r, &conf)
	return conf, err
}
