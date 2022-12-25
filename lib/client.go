package lib

import (
	"net/http"
	"net/url"
)

type Client struct {
	// HTTP client used for making requests.
	Http *http.Client

	// Base URL for API requests.
	BaseUrl *url.URL

	// User-agent used when communicating with the API.
	UserAgent string

	// The API key used for authentication & authorization.
	ApiKey string
}

func (c *Client) SetHttpClient(httpClient *http.Client) {
	c.Http = httpClient
}

func (c *Client) SetBaseUrl(baseUrlStr string) error {
	baseUrl, err := url.Parse(baseUrlStr)
	if err != nil {
		return err
	}

	c.BaseUrl = baseUrl

	return nil
}

func (c *Client) SetUserAgent(userAgent string) {
	c.UserAgent = userAgent
}

func (c *Client) SetApiKey(apiKey string) {
	c.ApiKey = apiKey
}
