package api

import (
	"count-jobs/collector"
	"net/http"
)

// example url with querys http://localhost:3000/api?location=Paris&term=PHP

// DataHandler handles data requests
func DataHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		location := r.URL.Query().Get("location")
		term := r.URL.Query().Get("term")
		country := r.URL.Query().Get("country")
		result := collector.StartCollector(term, location, country)
		w.Write([]byte(result))
	}
}
