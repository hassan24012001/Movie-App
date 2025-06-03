package movies_client

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	ApiKey  string
	BaseUrl string
}

func NewMoviesClient() *Client {
	return &Client{
		ApiKey:  apiKey,
		BaseUrl: baseUrl,
	}
}

// Since we want to reuse this method so return interface here
// Todo: Refactor response to return specific response map to specific endpoint
func (c *Client) GetResponse(url string, params interface{}) (interface{}, error) {
	req, err := http.NewRequest("GET", c.BaseUrl+url, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create http request")
	}

	query := req.URL.Query()
	// Set the API key
	query.Add("api_key", c.ApiKey)

	if params != nil {
		for key, value := range params.(map[string]interface{}) {
			query.Add(key, fmt.Sprintf("%v", value))
		}
	}

	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to make http request")
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Print("Error closing response body:", err)
		}
	}()

	var body bytes.Buffer
	if _, err := io.Copy(&body, resp.Body); err != nil {
		return nil, errors.Wrapf(err, "error reading response body")
	}

	log.Print("Get call successful")
	return body.Bytes(), nil
}
