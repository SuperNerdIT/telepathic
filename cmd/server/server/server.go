package server

import (
	"encoding/json"
	"net/http"

	"github.com/SuperNerdIT/telepathic/cmd/server/configs"
)

// this funtion could receive the configs
// and also cold receive the already initialized router/mux with all the handlers needed
// the function that initializes the hanlders and the routers could receive other things
// so the endpoints could make use of those dependencies.
func NewServer(c *configs.Configs) *http.Server {
	r := http.DefaultServeMux
	// refactor this, maybe another package for handlers ?
	r.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(struct{ Ok bool }{Ok: true})
	})
	server := http.Server{
		Addr:    c.Host + ":" + c.Port,
		Handler: r,
	}

	return &server
}
