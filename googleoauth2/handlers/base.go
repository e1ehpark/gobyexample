package handlers

import (
	"log"
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Handler start")
	mux.Handle("/login", http.FileServer(http.Dir("../templates")))
	log.Println("/ started")

	mux.HandleFunc("/auth/google/login", oauthGoogleLogin)
	mux.HandleFunc("/auth/google/callback", oauthGoogleCallback)

	return mux
}
