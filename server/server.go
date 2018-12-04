package server

import (
	"flag"

	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
)

func init() {
	log.Println("Initializing server")
}

//Start starts the server
func Start() {
	log.Info("Starting the Query")

	router := NewRouter()
	http.Handle("/", router)

	log.Infof("Server listening on port  %s \n", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
