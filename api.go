package main
import (
	"fmt"
	"os"

	"net/http"

	"github.com/peppage/foursquarego"
)

func main() {
	httpClient := http.DefaultClient
	// When creating the client you can specify either clientSecret or the accesstoken
	client := foursquarego.NewClient(httpClient, "foursquare", os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), "")
	fmt.Println(client)

	// Search Venues
	venues, _, _ := client.Venues.Search(&foursquarego.VenueSearchParams{
		LatLong:    "41.293209, -95.996140",
		Radius:     2000,
		CategoryID: []string{"4bf58dd8d48988d1ca941735"},
		Intent:     foursquarego.IntentCheckin,
	})
	for _, loc := range venues {
		fmt.Println(loc.Location)
	}

}

