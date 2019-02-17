package main
import (
	"fmt"
	"net/http"

	"github.com/peppage/foursquarego"
)

func main() {
	fmt.Println("Whoo")
	httpClient := http.DefaultClient
	// When creating the client you can specify either clientSecret or the accesstoken
	client := foursquarego.NewClient(httpClient, "foursquare", "clientId", "clientSecret", "")

}

