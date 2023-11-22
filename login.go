package ninjarmm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Login to the NinjaRMM API with valid `cliendID`, `clientSecret` and `scope`.
//
// See https://eu.ninjarmm.com/apidocs-beta/authorization/create-applications/machine-to-machine-apps to create your client ID and secret.
//
// Usage:
//
//	err := Login("clientID", "clientSecret", "monitoring management control")
//	if err != nil {
//		panic(err)
//	}
func Login(options ...string) (err error) {

	now := time.Now()

	// Check if we already have a valid token
	if auth != nil && auth.expiresAt.After(now) {
		return
	} else if auth == nil && len(options) != 3 {
		err = fmt.Errorf("error logging in, no client ID, secret or scope provided")
		return
	} else if auth != nil {
		options = []string{auth.clientID, auth.clientSecret, auth.Scope}
	}

	values := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {options[0]},
		"client_secret": {options[1]},
		"scope":         {options[2]},
	}

	req, err := http.NewRequest(http.MethodPost, authUrl, strings.NewReader(values.Encode()))
	if err != nil {
		err = fmt.Errorf("error creating request: %w", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("error sending request: %w", err)
		return
	}

	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusOK {
		body, e := io.ReadAll(res.Body)
		if err != nil {
			err = fmt.Errorf("error reading response body: %w", e)
			return
		}
		err = fmt.Errorf("error bad status code '%d': %s", status, body)
		return
	}

	var response authResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		err = fmt.Errorf("error decoding response body: %w", err)
	} else if response.AccessToken == "" || response.ExpiresIn == 0 {
		err = errors.New("no valid access token found in response")
	} else {
		response.expiresAt = now.Add(time.Duration(response.ExpiresIn-60) * time.Second)
		response.clientID = values["client_id"][0]
		response.clientSecret = values["client_secret"][0]
		auth = &response
	}

	return
}

// Response from the NinjaRMM API when logging in.
//
// It contains the access token and other information.
// See https://eu.ninjarmm.com/apidocs-beta/authorization/overview for more information.
type authResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`

	// Internal fields
	expiresAt    time.Time
	clientID     string
	clientSecret string
}
