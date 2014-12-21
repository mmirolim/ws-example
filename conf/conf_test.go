package conf

import (
	"fmt"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	// read file and check parsing
	f := `
[ds]
	[ds.redis]
	port = 6379
	chan = "orders"

[ws]
port = 12345
`
	// convert to Reader interface
	r := strings.NewReader(f)
	app, err := Read(r)
	if err != nil {
		t.Error("Read error")
	}
	fmt.Printf("%+v\n", app)
	want := 6379
	if got := app.DS.Redis.Port; got != want {
		t.Errorf("Datastore redis port %d, want %d", got, want)
	}
	want = 12345
	if got := app.WS.Port; got != want {
		t.Errorf("Websocket port %d, want %d", got, want)
	}

}
