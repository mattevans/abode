# abode 🏠 

[![GoDoc](https://godoc.org/github.com/mattevans/abode?status.svg)](https://godoc.org/github.com/mattevans/abode)
[![Build Status](https://travis-ci.org/mattevans/abode.svg?branch=master)](https://travis-ci.org/mattevans/abode)
[![Go Report Card](https://goreportcard.com/badge/github.com/mattevans/abode)](https://goreportcard.com/report/github.com/mattevans/abode)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/mattevans/abode/blob/master/LICENSE)

Explode one-line address strings using Golang. 

This package uses the [Google Maps API](https://console.developers.google.com/apis/credentials) to geocode the address.
Specifically you will require the [Geocoding API](https://console.developers.google.com/apis/library/geocoding-backend.googleapis.com)
enabled to translate address strings to detailed address objects.

Don't forget to set your `GOOGLE_MAPS_API_KEY` environment variable.

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

Configuration
-------------

Each `abode.Address{}` component can be tailored to meet your needs. Simply adjust the mapping of the Google Maps Address Components [here](https://github.com/mattevans/abode/blob/master/component.go#L31).


Disclaimer
-------------

Ensure your end results are used in conjunction with a Google Map to avoid violating the [Google Maps API Terms of Service](https://developers.google.com/maps/documentation/geocoding/policies).
