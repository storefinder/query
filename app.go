package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/storefinder/query/server"
)

func init() {
	log.Info("Initializing storelocator-query")
}

func main() {
	server.Start()
}
