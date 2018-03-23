package externalapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
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
	client := http.Client{Timeout: time.Millisecond * timeout}

	resp, err := client.Get(URI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, p)
	if err != nil {
		return err
	}

	return nil
}
