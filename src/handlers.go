package main

import (
	"log"
	"net/http"
)

func RootHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("root handler")
		w.WriteHeader(http.StatusOK)
	})
}

func PostsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("posts handler")
		w.WriteHeader(http.StatusOK)
	})
}