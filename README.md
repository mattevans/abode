# abode 🏠 

[![GoDoc](https://godoc.org/github.com/mattevans/abode?status.svg)](https://godoc.org/github.com/mattevans/abode)
[![Build Status](https://travis-ci.org/mattevans/abode.svg?branch=master)](https://travis-ci.org/mattevans/abode)
[![Go Report Card](https://goreportcard.com/badge/github.com/mattevans/abode)](https://goreportcard.com/report/github.com/mattevans/abode)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/mattevans/abode/blob/master/LICENSE)

- Geocode one-line addresses.
- Determine timezone information for a given address.
- This package uses the [Google Maps Web Services](https://developers.google.com/maps/web-services) to geocode the address.
- You will require the [Geocoding API](https://developers.google.com/maps/documentation) enabled, and optionally the [Timezone API] if you wish to also use `Timezone()`.
- Remember to set your `GOOGLE_MAPS_API_KEY` environment variable.

Installation
-----------------

`go get -u github.com/mattevans/abode`

Example
-------------

### Geocode an address:

```go
addr := "193 Rogers Ave, Brooklyn, New York"

address, err := abode.ExplodeWithContext(ctx, addr)
if err != nil {
  return err
}
```

Returns...

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

### Timezone information for an address:

```go
addr := "193 Rogers Ave, Brooklyn, New York"

address, err := abode.Timezone(ctx, addr)
if err != nil {
  return err
}
```

Returns...

```go
abode.Location{
    DstOffset:     0,
    RawOffset:     -17762,
    TimeZoneId:    "GMT-04:56:02",
    TimeZoneName:  "America/New_York"
}
```

Disclaimer
-------------

Ensure your end results are used in conjunction with a Google Map to avoid violating the [Google Maps API Terms of Service](https://developers.google.com/maps/documentation/geocoding/policies).
