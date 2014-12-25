package routes

import (
	"fmt"
	"io"
	"net/http"
)

func Initialize() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hi", sayHi)
	mux.HandleFunc("/last", getLastOrders)
	return mux
}

func sayHi(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hi\n")
}

func getLastOrders(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Request: %+v\n", req)
}
