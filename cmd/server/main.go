package main

import (
	"log"

	"github.com/SuperNerdIT/telepathic/configs"
	"github.com/SuperNerdIT/telepathic/server"
)

func main() {
	// Here we could make the configs first and then inject those into the server
	cfg := configs.MakeConfigs()
	srv := server.NewServer(cfg)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
