package api

import (
	"count-jobs/collector"
	"net/http"
)

// DataHandler handles data requests
func DataHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		collector.StartCollector()
	}
}
