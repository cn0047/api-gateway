package externalapi

import (
	"github.com/cnkint/curl"
	"time"
)

const (
	timeout = 10000 // milliseconds
)

// TakeData - performs request into external JSON REST API.
// This method performs request,
// than unmarshal JSON and assigns payload to provided pointer.
// It's possible to use this method for any external JSON REST API.
func TakeData(URI string, p interface{}) error {
	return curl.Unmarshal(curl.Options{URL: URI, Timeout: time.Millisecond * timeout}, p)
}
