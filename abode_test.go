package abode

import (
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	. "github.com/onsi/gomega"
)

func TestExplodeUSAddress(t *testing.T) {
	// Register the test.
	RegisterTestingT(t)

	// Init client
	if client == nil {
		err := initClient()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	}

	line1 := "193 Rogers Avenue"
	line2 := "Brooklyn"
	state := "New York"
	country := "United States"
	countryCode := "US"
	zip := "11216"
	lat := 40.6706153
	lng := -73.9530882
	formatted := "193 Rogers Ave, Brooklyn, NY 11216, USA"

	expected := &Address{
		AddressLine1:     &line1,
		AddressLine2:     &line2,
		AddressCity:      nil,
		AddressState:     &state,
		AddressCountry:   &country,
		AddressCountryCode: &countryCode,
		AddressZip:       &zip,
		AddressLat:       &lat,
		AddressLng:       &lng,
		FormattedAddress: &formatted,
	}

	address1 := "193 Rogers Ave, Brooklyn, New York"
	result, err := Explode(address1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected Results\n%s", pretty.Compare(result, expected))
	}
}

func TestExplodeInternationalAddress(t *testing.T) {
	// Register the test.
	RegisterTestingT(t)

	// Init client
	if client == nil {
		err := initClient()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	}

	line1 := "1 4 Abercrombie Street"
	line2 := "Howick"
	city := "Auckland"
	state := "Auckland"
	country := "New Zealand"
	countryCode := "NZ"
	zip := "2014"
	lat := -36.8991018
	lng := 174.9338525
	formatted := "1/4 Abercrombie St, Howick, Auckland 2014, New Zealand"

	expected := &Address{
		AddressLine1:     &line1,
		AddressLine2:     &line2,
		AddressCity:      &city,
		AddressState:     &state,
		AddressCountry:   &country,
		AddressCountryCode: &countryCode,
		AddressZip:       &zip,
		AddressLat:       &lat,
		AddressLng:       &lng,
		FormattedAddress: &formatted,
	}

	address1 := "1/4 Abercrombie Street, Auckland"
	result, err := Explode(address1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected Results\n%s", pretty.Compare(result, expected))
	}
}
