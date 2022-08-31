package abode

import (
	"context"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	. "github.com/onsi/gomega"
)

func TestExplode_US(t *testing.T) {
	RegisterTestingT(t)

	if client == nil {
		if err := mapsClient(); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	}

	var (
		line1       = "193 Rogers Avenue"
		line2       = "Brooklyn"
		state       = "New York"
		country     = "United States"
		countryCode = "US"
		zip         = "11216"
		lat         = 40.6706252
		lng         = -73.9530545
		formatted   = "193 Rogers Ave, Brooklyn, NY 11216, USA"

		expected = &Address{
			AddressLine1:       &line1,
			AddressLine2:       &line2,
			AddressCity:        nil,
			AddressState:       &state,
			AddressCountry:     &country,
			AddressCountryCode: &countryCode,
			AddressZip:         &zip,
			AddressLat:         &lat,
			AddressLng:         &lng,
			FormattedAddress:   &formatted,
		}
	)

	result, err := ExplodeWithContext(context.Background(), "193 Rogers Ave, Brooklyn, New York")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected Results\n%s", pretty.Compare(result, expected))
	}
}

func TestExplode_International(t *testing.T) {
	RegisterTestingT(t)

	if client == nil {
		if err := mapsClient(); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	}

	var (
		line1       = "1 4 Abercrombie Street"
		line2       = "Howick"
		city        = "Auckland"
		state       = "Auckland"
		country     = "New Zealand"
		countryCode = "NZ"
		zip         = "2014"
		lat         = -36.8990751
		lng         = 174.9334851
		formatted   = "1/4 Abercrombie Street, Howick, Auckland 2014, New Zealand"

		expected = &Address{
			AddressLine1:       &line1,
			AddressLine2:       &line2,
			AddressCity:        &city,
			AddressState:       &state,
			AddressCountry:     &country,
			AddressCountryCode: &countryCode,
			AddressZip:         &zip,
			AddressLat:         &lat,
			AddressLng:         &lng,
			FormattedAddress:   &formatted,
		}
	)

	result, err := ExplodeWithContext(context.Background(), "1/4 Abercrombie Street, Auckland")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected Results\n%s", pretty.Compare(result, expected))
	}
}

func TestTimezone_US(t *testing.T) {
	RegisterTestingT(t)

	if client == nil {
		if err := mapsClient(); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	}

	var (
		expected = &Location{
			DstOffset:    0,
			RawOffset:    -17762,
			TimeZoneId:   "America/New_York",
			TimeZoneName: "GMT-04:56:02",
		}
	)

	result, err := Timezone(context.Background(), "193 Rogers Ave, Brooklyn, New York")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected Results\n%s", pretty.Compare(result, expected))
	}
}

func TestTimezone_International(t *testing.T) {
	RegisterTestingT(t)

	if client == nil {
		if err := mapsClient(); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	}

	var (
		expected = &Location{
			DstOffset:    0,
			RawOffset:    41944,
			TimeZoneId:   "Pacific/Auckland",
			TimeZoneName: "GMT+11:39:04",
		}
	)

	result, err := Timezone(context.Background(), "1/4 Abercrombie Street, Auckland")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected Results\n%s", pretty.Compare(result, expected))
	}
}
