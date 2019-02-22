package pizza
import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	lat := q.Get("lat")
	lon := q.Get("lon")
	venuesJson, err := Venues(lat, lon)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, string(v))
}

