package main

import (
	"count-jobs/api"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello from Golang API 👋🏼"))
	}
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/api", api.DataHandler)
	http.ListenAndServe(":3000", nil)
}
