package email_verify

import (
	"net/http"

	"github.com/mslmio/sdk-go/lib"
)

var (
	DefaultClient     *Client
	DefaultHttpClient = http.DefaultClient
	DefaultBaseUrl    = "https://mslm.io"
	DefaultUserAgent  = "mslm/go/1.1.1"
	DefaultApiKey     = ""
)

type Client struct {
	// Common client system.
	C *lib.Client
}

func Init(
	apiKey string,
) *Client {
	c := &Client{}
	c.C = new(lib.Client)
	c.SetHttpClient(http.DefaultClient)
	c.SetBaseUrl(DefaultBaseUrl)
	c.SetUserAgent(DefaultUserAgent)
	c.SetApiKey(apiKey)
	return c
}

func InitDefaults() *Client {
	return Init(DefaultApiKey)
}

func (c *Client) SetHttpClient(httpClient *http.Client) {
	c.C.SetHttpClient(httpClient)
}

func SetHttpClient(httpClient *http.Client) {
	DefaultClient.SetHttpClient(httpClient)
}

func (c *Client) SetBaseUrl(baseUrlStr string) error {
	return c.C.SetBaseUrl(baseUrlStr)
}

func SetBaseUrl(baseUrlStr string) error {
	return DefaultClient.SetBaseUrl(baseUrlStr)
}

func (c *Client) SetUserAgent(userAgent string) {
	c.C.SetUserAgent(userAgent)
}

func SetUserAgent(userAgent string) {
	DefaultClient.SetUserAgent(userAgent)
}

func (c *Client) SetApiKey(apiKey string) {
	c.C.SetApiKey(apiKey)
}

func SetApiKey(apiKey string) {
	DefaultClient.SetApiKey(apiKey)
}
