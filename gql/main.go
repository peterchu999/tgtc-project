package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/peterchu999/tgtc-project/gql/gqlserver"
	"github.com/peterchu999/tgtc-project/gql/server"
)

func main() {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("gql/webstatic"))
	router.PathPrefix("/").Handler(fs)

	handler := gqlserver.Handle()
	router.Path("/graphql").Handler(handler)
	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         9000,
	}
	server.Serve(serverConfig, router)
}