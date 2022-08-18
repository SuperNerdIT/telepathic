package server

import "log"

func StartServer() {
	srv := NewServer()
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
