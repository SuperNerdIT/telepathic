package server

import (
	"fmt"
	"net/http"
)


func NewServer() *http.Server {
	r := http.DefaultServeMux
	r.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Ok")

	})
	server := http.Server{
		Addr: "localhost:3000",
		Handler: r,

	}
	
	
	return &server
}