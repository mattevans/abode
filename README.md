# abode 🏠 

[![GoDoc](https://godoc.org/github.com/mattevans/abode?status.svg)](https://godoc.org/github.com/mattevans/abode)
[![Build Status](https://travis-ci.org/mattevans/abode.svg?branch=master)](https://travis-ci.org/mattevans/abode)

Explode one-line address strings using Golang. 

This package uses the [Google Maps API](https://console.developers.google.com/apis/credentials) to geocode the address. Don't forget to set your `GOOGLE_MAPS_API_KEY` environment variable.

Installation
-----------------

`go get -u github.com/mattevans/abode`

Example
-------------

Explode your one-line address...

```go
yourAddress := "193 Rogers Ave, Brooklyn, New York"

// Explode our one-line address into components.
address, err := abode.Explode(yourAddress)
if err != nil {
  return err
}
```

Which will give you...

```go
abode.Address{
  AddressLine1:     "193 Rogers Avenue",
  AddressLine2:     "Brooklyn"
  AddressCity:      nil,
  AddressState:     "New York"
  AddressCountry:   "United States"
  AddressZip:       "11216"
  AddressLat:       40.6706073,
  AddressLng:       -73.9530182,
  FormattedAddress: "193 Rogers Ave, Brooklyn, NY 11216, USA",
}
```

Ensure your end results are used in conjunction with a Google Map to avoid violating the [Google Maps API Terms of Service](https://developers.google.com/maps/documentation/geocoding/policies).
