package main
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	pizza "github.com/mattdsteele/pizza-compass"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		lat := q.Get("lat")
		lon := q.Get("lon")
		fmt.Println("lat", lat)
		fmt.Println("lon", lon)
		// lat := "41.293209"
		// lon := "-95.996140"
		venues := pizza.Venues(lat, lon)
		v, err := json.Marshal(venues)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, string(v))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}

