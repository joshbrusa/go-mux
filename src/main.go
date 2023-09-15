package main

import (
	"net/http"
)

func main() {
	port := "8000"
	mux := http.NewServeMux()

	server := Server{
		Port: port,
		Mux: mux,
	}

	server.Start()
}