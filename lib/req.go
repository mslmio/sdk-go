package lib

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type ReqOpts struct {
	// Override the configured HTTP client just for this request.
	Http *http.Client

	// Override the configured base URL just for this request.
	BaseUrl *url.URL

	// Override the configured user-agent just for this request.
	UserAgent *string

	// Override the configured API key just for this request.
	ApiKey *string

	// Context to use for this request.
	//
	// Defaults to `context.Background()`.
	Context context.Context
}

// Returns a version of ReqOpts coalesced with documented & client defaults.
func (c *Client) PrepareOpts(opt *ReqOpts) ReqOpts {
	if opt == nil {
		return ReqOpts{}
	}

	// Prepare http client.
	httpC := c.Http
	if opt.Http != nil {
		httpC = opt.Http
	}

	// Prepare base URL.
	baseUrl := c.BaseUrl
	if opt.BaseUrl != nil {
		baseUrl = opt.BaseUrl
	}

	// Prepare user-agent.
	userAgent := c.UserAgent
	if opt.UserAgent != nil {
		userAgent = *opt.UserAgent
	}

	// Prepare API key.
	apiKey := c.ApiKey
	if opt.ApiKey != nil {
		apiKey = *opt.ApiKey
	}

	// Prepare context.
	ctx := opt.Context
	if opt.Context == nil {
		ctx = context.Background()
	}

	return ReqOpts{
		Http:      httpC,
		BaseUrl:   baseUrl,
		UserAgent: &userAgent,
		ApiKey:    &apiKey,
		Context:   ctx,
	}
}

func (c *Client) PrepareUrl(
	urlPath string,
	qp map[string]string,
	opt *ReqOpts,
) (*url.URL, error) {
	tUrl, err := opt.BaseUrl.Parse(urlPath)
	if err != nil {
		return nil, err
	}

	tUrlQp := url.Values{}
	for k, v := range qp {
		tUrlQp.Add(k, v)
	}
	tUrlQp.Set("apikey", *opt.ApiKey)
	tUrl.RawQuery = tUrlQp.Encode()

	return tUrl, nil
}

func (c *Client) ReqAndResp(
	method string,
	tUrl *url.URL,
	respData interface{},
	opt *ReqOpts,
) error {
	req, err := http.NewRequestWithContext(
		opt.Context,
		method,
		tUrl.String(),
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", *opt.UserAgent)

	// Do request.
	resp, err := opt.Http.Do(req)
	if err != nil {
		return err
	}

	// Read body.
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Decode request.
	if err := json.Unmarshal(body, respData); err != nil {
		return err
	}

	return nil
}
