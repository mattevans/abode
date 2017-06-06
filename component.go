package abode

import "googlemaps.github.io/maps"

// Define the different address component types.
const (
	AddressComponentTypeSubPremise      = "subpremise"
	AddressComponentTypePremise         = "premise"
	AddressComponentTypeStreetNumber    = "street_number"
	AddressComponentTypeRoute           = "route"
	AddressComponentTypeStreetAddress   = "street_address"
	AddressComponentTypeSubLocality     = "sublocality"
	AddressComponentTypeLocality        = "locality"
	AddressComponentTypeAdminAreaLevel1 = "administrative_area_level_1"
	AddressComponentTypeAdminAreaLevel2 = "administrative_area_level_2"
	AddressComponentTypeAdminAreaLevel3 = "administrative_area_level_3"
	AddressComponentTypeAdminAreaLevel4 = "administrative_area_level_4"
	AddressComponentTypeAdminAreaLevel5 = "administrative_area_level_5"
	AddressComponentTypeCountry         = "country"
	AddressComponentTypePostalCode      = "postal_code"
)

// Defines the address components that make up Address.AddressLine1.
var addressLine1Composition = []string{
	AddressComponentTypeSubPremise,
	AddressComponentTypePremise,
	AddressComponentTypeStreetNumber,
	AddressComponentTypeRoute,
	AddressComponentTypeStreetAddress,
}

// Defines the address components that make up Address.AddressLine2.
var addressLine2Composition = []string{
	AddressComponentTypeSubLocality,
}

// Defines the address components that make up Address.AddressCity.
var addressCityComposition = []string{
	AddressComponentTypeLocality,
}

// Defines the address components that make up Address.AddressState.
var addressStateComposition = []string{
	AddressComponentTypeAdminAreaLevel1,
}

// Defines the address components that make up Address.Country.
var addressCountryComposition = []string{
	AddressComponentTypeCountry,
}

// Defines the address components that make up Address.CountryCode.
var addressCountryCodeComposition = []string{
	AddressComponentTypeCountry,
}

// Defines the address components that make up Address.Zip.
var addressPostalCodeComposition = []string{
	AddressComponentTypePostalCode,
}

// getComponentByType will return the given componentType from a slice of
// maps.AddressComponent's.
func getComponentByType(components []maps.AddressComponent, componentType string) *maps.AddressComponent {
	for k, component := range components {
		if isComponentType(component.Types, componentType) {
			return &components[k]
		}
	}
	return nil
}

// isComponentType will check to see if e is contained within s.
func isComponentType(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
