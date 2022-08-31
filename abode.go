package abode

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"googlemaps.github.io/maps"
)

var (
	client               *maps.Client
	addressLineDelimiter = ","
)

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

type Location struct {
	DstOffset    int    `json:"dst_offset"`
	RawOffset    int    `json:"raw_offset"`
	TimeZoneId   string `json:"time_zone_id"`
	TimeZoneName string `json:"time_zone_name"`
}

// client will initialize a Google Maps API client.
func mapsClient() error {
	var (
		err error
		key = os.Getenv("GOOGLE_MAPS_API_KEY")
	)

	if key == "" {
		return errors.New("missing `GOOGLE_MAPS_API_KEY`")
	}

	client, err = maps.NewClient(maps.WithAPIKey(key))
	return err
}

// Explode takes a one-line address string, explodes it using gmaps.Geocode() and returns an *Address.
// Deprecated: Use ExplodeWithContext instead.
func Explode(address string) (*Address, error) {
	return ExplodeWithContext(context.Background(), address)
}

// ExplodeWithContext takes a one-line address string, explodes it using gmaps.Geocode() and returns an *Address.
func ExplodeWithContext(ctx context.Context, address string) (*Address, error) {
	if client == nil {
		if err := mapsClient(); err != nil {
			return nil, err
		}
	}

	rsp, err := client.Geocode(ctx, &maps.GeocodingRequest{
		Address: address,
	})
	if err != nil {
		return nil, err
	}

	if len(rsp) < 1 {
		return nil, err
	}

	var (
		geocodeResult    = rsp[0]
		components       = geocodeResult.AddressComponents
		formattedAddress = geocodeResult.FormattedAddress
		lat              = geocodeResult.Geometry.Location.Lat
		lng              = geocodeResult.Geometry.Location.Lng
	)

	return &Address{
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
	}, err
}

// Timezone takes a one-line address string, and determine timezone/location data for it using gmaps.Timezone().
func Timezone(ctx context.Context, address string) (*Location, error) {
	if client == nil {
		if err := mapsClient(); err != nil {
			return nil, err
		}
	}

	geocodeResult, err := ExplodeWithContext(ctx, address)
	if err != nil {
		return nil, err
	}

	if geocodeResult.AddressLat == nil || geocodeResult.AddressLng == nil {
		return nil, errors.New("unable to determine latitude and longitude for address")
	}

	resp, err := client.Timezone(ctx, &maps.TimezoneRequest{
		Location: &maps.LatLng{Lat: *geocodeResult.AddressLat, Lng: *geocodeResult.AddressLng},
	})
	if err != nil {
		return nil, err
	}

	return &Location{
		DstOffset:    resp.DstOffset,
		RawOffset:    resp.RawOffset,
		TimeZoneId:   resp.TimeZoneID,
		TimeZoneName: resp.TimeZoneName,
	}, err
}

func compose(composition []string, delimiter string, components []maps.AddressComponent, useShortName bool) *string {
	var str string

	for _, element := range composition {
		component := getComponentByType(components, element)
		if useShortName {
			if component != nil && !strings.Contains(str, component.ShortName) {
				str = fmt.Sprintf("%s %s%s", str, component.ShortName, delimiter)
			}

			continue
		}

		if component != nil && !strings.Contains(str, component.LongName) {
			str = fmt.Sprintf("%s %s%s", str, component.LongName, delimiter)
		}
	}

	if str == "" {
		return nil
	}

	str = strings.TrimPrefix(strings.TrimSuffix(str, delimiter), " ")

	return &str
}
