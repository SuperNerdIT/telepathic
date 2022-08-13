package server

import (
	"encoding/json"
	"net/http"
)


func NewServer() *http.Server {
	r := http.DefaultServeMux
	r.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(struct{Ok bool}{Ok: true})

	})
	server := http.Server{
		Addr: "localhost:3000",
		Handler: r,
		

	}
	
	
	return &server
}