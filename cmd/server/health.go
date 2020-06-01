package main

import "net/http"

var healthURI = "/health"

func health(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("ok"))
}