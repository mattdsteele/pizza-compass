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
		Intent:     foursquarego.IntentCheckin,
	})

	if err != nil {
		return nil, -1, -1, err
	}
	sort.Slice(venues, func(i, j int) bool {
		iDist := Distance(venues[i].Location.Lat, venues[i].Location.Lng, userLat, userLon)
		jDist := Distance(venues[j].Location.Lat, venues[j].Location.Lng, userLat, userLon)
		return iDist < jDist
	})
	rateLimit := foursquarego.ParseRate(resp)
	fmt.Println(venues, err)
	fmt.Println("Have " + strconv.Itoa(rateLimit.Remaining) + " queries left")
	return venues, rateLimit.Limit, rateLimit.Remaining, nil
}

