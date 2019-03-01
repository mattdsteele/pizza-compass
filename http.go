package pizza
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	q := r.URL.Query()
	lat := q.Get("lat")
	lon := q.Get("lon")
	venuesJson, limit, limitLeft, err := Venues(lat, lon)
	w.Header().Set("X-API-Limit", strconv.Itoa(limit))
	w.Header().Set("X-API-Limit-Left", strconv.Itoa(limitLeft))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	v, err := json.Marshal(venuesJson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(v))
}

