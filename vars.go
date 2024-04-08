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

// Base request used by all other requests
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

	req.Header.Set("Authorization", "Bearer "+auth.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("error sending request: %w", err)
		return
	}

	defer res.Body.Close()

	if status := res.StatusCode; status > 299 {
		body, _ := io.ReadAll(res.Body)
		err = fmt.Errorf("error bad status code '%d' : %s", status, body)
		return
	}

	if response != nil {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			err = fmt.Errorf("error decoding response body: %w", err)
		}
	}

	return
}

// Base request for multipart/form-data requests
// func UploadMultipartFile(client *http.Client, uri, key, path string) (*http.Response, error) {
// 	body, writer := io.Pipe()

// 	req, err := http.NewRequest(http.MethodPost, uri, body)
// 	if err != nil {
// 			return nil, err
// 	}

// 	mwriter := multipart.NewWriter(writer)
// 	req.Header.Add("Content-Type", mwriter.FormDataContentType())

// 	errchan := make(chan error)

// 	go func() {
// 			defer close(errchan)
// 			defer writer.Close()
// 			defer mwriter.Close()

// 			w, err := mwriter.CreateFormFile(key, path)
// 			if err != nil {
// 					errchan <- err
// 					return
// 			}

// 			in, err := os.Open(path)
// 			if err != nil {
// 					errchan <- err
// 					return
// 			}
// 			defer in.Close()

// 			if written, err := io.Copy(w, in); err != nil {
// 					errchan <- fmt.Errorf("error copying %s (%d bytes written): %v", path, written, err)
// 					return
// 			}

// 			if err := mwriter.Close(); err != nil {
// 					errchan <- err
// 					return
// 			}
// 	}()

// 	resp, err := client.Do(req)
// 	merr := <-errchan

// 	if err != nil || merr != nil {
// 			return resp, fmt.Errorf("http error: %v, multipart error: %v", err, merr)
// 	}

// 	return resp, nil
// }
