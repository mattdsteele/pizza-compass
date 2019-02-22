package pizza
import (
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/peppage/foursquarego"
)

func Venues(lat, lon string) []foursquarego.Venue {
	userLat, _ := strconv.ParseFloat(lat, 64)
	userLon, _ := strconv.ParseFloat(lon, 64)
	httpClient := http.DefaultClient
	// When creating the client you can specify either clientSecret or the accesstoken
	client := foursquarego.NewClient(httpClient, "foursquare", os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), "")

	// Search Venues
	searchParam := lat + ", " + lon
	pizzaCat := "4bf58dd8d48988d1ca941735"
	venues, _, _ := client.Venues.Search(&foursquarego.VenueSearchParams{
		LatLong: searchParam,
		// Radius:     2000,
		CategoryID: []string{pizzaCat},
		Intent:     foursquarego.IntentCheckin,
	})
	sort.Slice(venues, func(i, j int) bool {
		iDist := Distance(venues[i].Location.Lat, venues[i].Location.Lng, userLat, userLon)
		jDist := Distance(venues[j].Location.Lat, venues[j].Location.Lng, userLat, userLon)
		return iDist < jDist
	})
	return venues

}

