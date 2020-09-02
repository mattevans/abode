你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# abode 🏠 

[![GoDoc](https://godoc.org/github.com/mattevans/abode?status.svg)](https://godoc.org/github.com/mattevans/abode)
[![Build Status](https://travis-ci.org/mattevans/abode.svg?branch=master)](https://travis-ci.org/mattevans/abode)
[![Go Report Card](https://goreportcard.com/badge/github.com/mattevans/abode)](https://goreportcard.com/report/github.com/mattevans/abode)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/mattevans/abode/blob/master/LICENSE)

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

Configuration
-------------

Each `abode.Address{}` component can be tailored to meet your needs. Simply adjust the mapping of the Google Maps Address Components [here](https://github.com/mattevans/abode/blob/master/component.go#L31).


Disclaimer
-------------

Ensure your end results are used in conjunction with a Google Map to avoid violating the [Google Maps API Terms of Service](https://developers.google.com/maps/documentation/geocoding/policies).
