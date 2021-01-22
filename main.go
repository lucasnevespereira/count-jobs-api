package main

import (
	"count-jobs/api"
	"net/http"
	"os"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello from the Count Jobs API 👋🏼"))
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/api", api.DataHandler)
	http.ListenAndServe(":"+port, nil)
}
