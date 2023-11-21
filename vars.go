package ninjarmm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	baseUrl string = "https://eu.ninjarmm.com"
	authUrl string = baseUrl + "/ws/oauth/token"
	apiUrl  string = baseUrl + "/v2/"
)

var (
	auth   *authResponse
	client *http.Client = &http.Client{
		Timeout: time.Minute,
	}
)

// request sends a request to the NinjaRMM API
func request(method, path string, payload interface{}, response interface{}) (err error) {

	// Check if we already have a valid token
	err = Login()
	if err != nil {
		err = fmt.Errorf("error logging in: %w", err)
		return
	}

	buffer := new(bytes.Buffer)
	if payload != nil {
		err = json.NewEncoder(buffer).Encode(payload)
		if err != nil {
			err = fmt.Errorf("error encoding request body: %w", err)
			return
		}
	}

	req, err := http.NewRequest(method, apiUrl+path, buffer)
	if err != nil {
		err = fmt.Errorf("error creating request: %w", err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+auth.AccessToken)
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("error sending request: %w", err)
		return
	}

	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		err = fmt.Errorf("error bad status code '%d' : %s", status, body)
	}

	if response != nil {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			err = fmt.Errorf("error decoding response body: %w", err)
		}
	}

	return
}
