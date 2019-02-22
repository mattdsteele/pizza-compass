package main
import (
	"log"
	"net/http"

	pizza "github.com/mattdsteele/pizza-compass"
)

func main() {
	http.HandleFunc("/", pizza.HttpHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

