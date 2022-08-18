package main

import (
	"log"

	"github.com/SuperNerdIT/telepathic/cmd/server/configs"
	"github.com/SuperNerdIT/telepathic/cmd/server/server"
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
