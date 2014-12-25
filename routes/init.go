package routes

import (
	"fmt"
	"io"
	"net/http"
)

func SayHi(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hi\n")
}

func GetLastOrders(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Request: %+v\n", req)
}
