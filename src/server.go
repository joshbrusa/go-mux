package main

import (
	"log"
	"net/http"
)

type Server struct {
	Port string
	Mux *http.ServeMux
}

func (server *Server) Start() {
	server.Mux.Handle("/", LoggingMiddleware(RootHandler()))
	server.Mux.Handle("/posts", LoggingMiddleware(AuthMiddleware(PostsHandler())))

	log.Println("server is listening on port: " + server.Port)
	log.Fatal(http.ListenAndServe(":" + server.Port, server.Mux))
}
