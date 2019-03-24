package pizza
import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/peppage/foursquarego"
)

func Venues(lat, lon string) ([]foursquarego.Venue, int, int, error) {
	userLat, _ := strconv.ParseFloat(lat, 64)
	userLon, _ := strconv.ParseFloat(lon, 64)
	httpClient := http.DefaultClient
	// When creating the client you can specify either clientSecret or the accesstoken
	client := foursquarego.NewClient(httpClient, "foursquare", os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), "")

	// Search Venues
	searchParam := lat + ", " + lon
	pizzaCat := "4bf58dd8d48988d1ca941735"
	venues, resp, err := client.Venues.Search(&foursquarego.VenueSearchParams{
		LatLong:    searchParam,
		CategoryID: []string{pizzaCat},
		Radius:     8000,
		Intent:     foursquarego.IntentBrowse,
	})

	if err != nil {
		return nil, -1, -1, err
	}
	pizzaVenues := []foursquarego.Venue{}
	for _, venue := range venues {
		for _, cat := range venue.Categories {
			if cat.ID == pizzaCat {
				pizzaVenues = append(pizzaVenues, venue)
			}
		}
	}
	sort.Slice(pizzaVenues, func(i, j int) bool {
		iDist := Distance(pizzaVenues[i].Location.Lat, pizzaVenues[i].Location.Lng, userLat, userLon)
		jDist := Distance(pizzaVenues[j].Location.Lat, pizzaVenues[j].Location.Lng, userLat, userLon)
		return iDist < jDist
	})
	rateLimit := foursquarego.ParseRate(resp)
	fmt.Println(venues, err)
	fmt.Println("Have " + strconv.Itoa(rateLimit.Remaining) + " queries left")
	return pizzaVenues, rateLimit.Limit, rateLimit.Remaining, nil
}

