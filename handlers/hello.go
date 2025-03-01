package handlers

import "net/http"

func Hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello"))
}
