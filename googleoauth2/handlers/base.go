package handlers

import (
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("templates/")))

	mux.handleFunc("/auth/google/login", oauthGoogleLogin)
	mux.handleFunc("/auth/google/callback", oauthGoogleCallback)

	return mux
}
