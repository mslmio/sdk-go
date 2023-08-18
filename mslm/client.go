package mslm

import (
	"net/http"

	"github.com/mslmio/sdk-go/email_verify"
	"github.com/mslmio/sdk-go/lib"
	"github.com/mslmio/sdk-go/otp"
)

var (
	DefaultClient     *Client
	DefaultHttpClient = http.DefaultClient
	DefaultBaseUrl    = "https://mslm.io"
	DefaultUserAgent  = "mslm/go/1.0.0"
	DefaultApiKey     = ""
)

type Client struct {
	// Common client system.
	C *lib.Client

	// The Email Verify API client.
	EmailVerify *email_verify.Client

	// The OTP client.
	Otp *otp.Client
}

func Init(
	apiKey string,
) *Client {
	c := &Client{}
	c.C = new(lib.Client)
	c.EmailVerify = email_verify.Init(apiKey)
	c.Otp = otp.Init(DefaultApiKey)
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
	c.EmailVerify.SetHttpClient(httpClient)
	c.Otp.SetHttpClient(httpClient)
}

func SetHttpClient(httpClient *http.Client) {
	DefaultClient.SetHttpClient(httpClient)
}

func (c *Client) SetBaseUrl(baseUrlStr string) error {
	if err := c.C.SetBaseUrl(baseUrlStr); err != nil {
		return err
	}

	if err := c.EmailVerify.SetBaseUrl(baseUrlStr); err != nil {
		return err
	}

	if err := c.Otp.SetBaseUrl(baseUrlStr); err != nil {
		return err
	}

	return nil
}

func SetBaseUrl(baseUrlStr string) error {
	return DefaultClient.SetBaseUrl(baseUrlStr)
}

func (c *Client) SetUserAgent(userAgent string) {
	c.C.SetUserAgent(userAgent)
	c.EmailVerify.SetUserAgent(userAgent)
	c.Otp.SetUserAgent(userAgent)
}

func SetUserAgent(userAgent string) {
	DefaultClient.SetUserAgent(userAgent)
}

func (c *Client) SetApiKey(apiKey string) {
	c.C.SetApiKey(apiKey)
	c.EmailVerify.SetApiKey(apiKey)
	c.Otp.SetApiKey(apiKey)
}

func SetApiKey(apiKey string) {
	DefaultClient.SetApiKey(apiKey)
}
