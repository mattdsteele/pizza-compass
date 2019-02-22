package main
import (
	"fmt"

	pizza "github.com/mattdsteele/pizza-compass"
)

func main() {
	lat := "41.293209"
	lon := "-95.996140"
	venues := pizza.Venues(lat, lon)
	for _, v := range venues {
		fmt.Println(v.Name)
	}
}

