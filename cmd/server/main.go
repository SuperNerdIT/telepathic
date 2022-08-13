package server

import "log"

func StartServer() {
	log.Fatal(NewServer().ListenAndServe())
}
