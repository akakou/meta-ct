package metact

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func fetchAPI[T any](method string, _url string, data url.Values) (*T, error) {
	var body io.Reader
	result := new(T)

	// parse URL
	u, err := url.Parse(_url)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %s", err)
	}

	// set the parameters
	if method == "GET" {
		u.RawQuery = data.Encode()
		body = bytes.NewReader([]byte{})

	} else if method == "POST" {
		encoded := []byte(data.Encode())
		body = bytes.NewReader(encoded)
	}

	// make request
	req, err := http.NewRequest(method, u.String(), body)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err)
	}

	// check status code
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error response: %s", resp.Status)
	}

	// read and parse response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err)
	}

	// fmt.Printf("resp body", string(respBody))

	resp.Body.Close()

	err = json.Unmarshal(respBody, result)

	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %s", err)
	}

	return result, nil
}
