package abode

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"googlemaps.github.io/maps"
)

var client *maps.Client
var addressLineDelimiter = ","

// Address represents a response Address from abode.
type Address struct {
	AddressLine1       *string  `json:"address_line1"`
	AddressLine2       *string  `json:"address_line2"`
	AddressCity        *string  `json:"address_city"`
	AddressState       *string  `json:"address_state"`
	AddressCountry     *string  `json:"address_country"`
	AddressCountryCode *string  `json:"address_country_code"`
	AddressZip         *string  `json:"address_zip"`
	AddressLat         *float64 `json:"address_lat"`
	AddressLng         *float64 `json:"address_lng"`
	FormattedAddress   *string  `json:"formatted_address"`
}

// initClient will initalize a Google Maps API client.
func initClient() error {
	var err error
	key := os.Getenv("GOOGLE_MAPS_API_KEY")
	if key == "" {
		return errors.New("please configure a `GOOGLE_MAPS_API_KEY`")
	}
	client, err = maps.NewClient(maps.WithAPIKey(key))
	return err
}

// Explode takes a one-line address string, explodes it and returns an *Address
func Explode(address string) (*Address, error) {
	if client == nil {
		if err := initClient(); err != nil {
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

	// Using the first match in our response, grab the values we need.
	components := resp[0].AddressComponents
	formattedAddress := resp[0].FormattedAddress
	lat := resp[0].Geometry.Location.Lat
	lng := resp[0].Geometry.Location.Lng

	// Construct the return *Address{}
	response := &Address{
		AddressLine1:       compose(addressLine1Composition, "", components, false),
		AddressLine2:       compose(addressLine2Composition, addressLineDelimiter, components, false),
		AddressCity:        compose(addressCityComposition, addressLineDelimiter, components, false),
		AddressState:       compose(addressStateComposition, addressLineDelimiter, components, false),
		AddressCountry:     compose(addressCountryComposition, addressLineDelimiter, components, false),
		AddressCountryCode: compose(addressCountryCodeComposition, addressLineDelimiter, components, true),
		AddressZip:         compose(addressPostalCodeComposition, addressLineDelimiter, components, false),
		AddressLat:         &lat,
		AddressLng:         &lng,
		FormattedAddress:   &formattedAddress,
	}

	return response, err
}

func compose(composition []string, delimiter string, components []maps.AddressComponent, useShortName bool) *string {
	var str string
	for _, element := range composition {
		component := getComponentByType(components, element)
		if useShortName {
			if component != nil && !strings.Contains(str, component.ShortName) {
				str = fmt.Sprintf("%s %s%s", str, component.ShortName, delimiter)
			}
		} else {
			if component != nil && !strings.Contains(str, component.LongName) {
				str = fmt.Sprintf("%s %s%s", str, component.LongName, delimiter)
			}
		}
	}
	if str == "" {
		return nil
	}
	str = strings.TrimPrefix(strings.TrimSuffix(str, delimiter), " ")
	return &str
}
