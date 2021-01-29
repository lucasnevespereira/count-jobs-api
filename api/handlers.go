package api

import (
	"count-jobs/collector"
	"net/http"
)

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }

// DataHandler handles data requests
func DataHandler(w http.ResponseWriter, r *http.Request) {
	// enableCors(&w)
	switch r.Method {
	case http.MethodGet:
		location := r.URL.Query().Get("location")
		term := r.URL.Query().Get("term")
		country := r.URL.Query().Get("country")
		result := collector.StartCollector(term, location, country)
		w.Write([]byte(result))
	}
}
