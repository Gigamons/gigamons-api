package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	gmux "github.com/gorilla/mux"
)

var mux = gmux.NewRouter()

// GetMux returns the MuxRouter
func GetMux() *gmux.Router {
	return mux
}

// StartServer starts the HTTP Server
func StartServer(hostname string, port int16) {
	fmt.Printf("API Is listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(hostname+":"+strconv.Itoa(int(port)), mux))
}
