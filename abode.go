package abode

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/snikch/api/config"

	"googlemaps.github.io/maps"
)

var client *maps.Client
var addressLineDelimeter = ","

// Address represents a response Address from abode.
type Address struct {
	AddressLine1     *string  `json:"address_line1"`
	AddressLine2     *string  `json:"address_line2"`
	AddressCity      *string  `json:"address_city"`
	AddressState     *string  `json:"address_state"`
	AddressCountry   *string  `json:"address_country"`
	AddressZip       *string  `json:"address_zip"`
	AddressLat       *float64 `json:"address_lat"`
	AddressLng       *float64 `json:"address_lng"`
	FormattedAddress *string  `json:"formatted_address"`
}

// initClient will initalize a Google Maps API client.
func initClient() error {
	var err error
	key := config.PrivateString("GOOGLE_MAPS_API_KEY")
	if key == "" {
		return errors.New("Please configure a `GOOGLE_MAPS_API_KEY`")
	}
	client, err = maps.NewClient(maps.WithAPIKey(key))
	return err
}

// Explode takes a one-line address string, explodes it and returns an *Address
func Explode(address string) (*Address, error) {
	if client == nil {
		err := initClient()
		if err != nil {
			return nil, err
		}
	}

	// Build the API request.
	req := &maps.GeocodingRequest{
		Address: address,
	}

	// Execute the request.
	resp, err := client.Geocode(context.Background(), req)
	if len(resp) < 1 {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	// Using the first/closest match in our response, grab the values we need.
	components := resp[0].AddressComponents
	formattedAddress := resp[0].FormattedAddress
	lat := resp[0].Geometry.Location.Lat
	lng := resp[0].Geometry.Location.Lng

	// Construct the return *Address{}
	response := &Address{
		AddressLine1:     compose(addressLine1Composition, "", components),
		AddressLine2:     compose(addressLine2Composition, addressLineDelimeter, components),
		AddressCity:      compose(addressCityComposition, addressLineDelimeter, components),
		AddressState:     compose(addressStateComposition, addressLineDelimeter, components),
		AddressCountry:   compose(addressCountryComposition, addressLineDelimeter, components),
		AddressZip:       compose(addressPostalCodeComposition, addressLineDelimeter, components),
		AddressLat:       &lat,
		AddressLng:       &lng,
		FormattedAddress: &formattedAddress,
	}

	return response, err
}

func compose(composition []string, delimeter string, components []maps.AddressComponent) *string {
	var str string
	for _, element := range composition {
		component := getComponentByType(components, element)
		if component != nil && !strings.Contains(str, component.LongName) {
			str = fmt.Sprintf("%s %s%s", str, component.LongName, delimeter)
		}
	}
	if str == "" {
		return nil
	}
	str = strings.TrimPrefix(strings.TrimSuffix(str, delimeter), " ")
	return &str
}
