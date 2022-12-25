package email_verify

import (
	"net/url"

	"github.com/mslmdev/sdk-go/lib"
)

type SingleVerifyReqOpts struct {
	// Common request options.
	lib.ReqOpts

	// If this is set to `true`, the email won't be URL encoded just for this
	// request.
	//
	// By default, the input email address is URL encoded.
	DisableUrlEncode *bool
}

type SingleVerifyResp struct {
	Email      string                `json:"email"`
	Username   string                `json:"username"`
	Domain     string                `json:"domain"`
	Malformed  bool                  `json:"malformed"`
	Suggestion string                `json:"suggestion"`
	Status     string                `json:"status"`
	HasMailbox bool                  `json:"has_mailbox"`
	AcceptAll  bool                  `json:"accept_all"`
	Disposable bool                  `json:"disposable"`
	Free       bool                  `json:"free"`
	Role       bool                  `json:"role"`
	Mx         []*SingleVerifyRespMx `json:"mx"`
}

type SingleVerifyRespMx struct {
	Host string `json:"host"`
	Pref int    `json:"pref"`
}

func (c *Client) SingleVerify(
	emailAddr string,
	opts ...*SingleVerifyReqOpts,
) (*SingleVerifyResp, error) {
	// Prepare options.
	opt := &SingleVerifyReqOpts{}
	if len(opts) > 0 {
		opt = opts[len(opts)-1]
	}
	opt.ReqOpts = c.C.PrepareOpts(&opt.ReqOpts)

	// Prepare URL.
	if opt.DisableUrlEncode != nil && !*opt.DisableUrlEncode {
		emailAddr = url.QueryEscape(emailAddr)
	}
	qp := map[string]string{"email": emailAddr}
	tUrl, err := c.C.PrepareUrl("/api/sv/v1", qp, &opt.ReqOpts)
	if err != nil {
		return nil, err
	}

	// Get data.
	svResp := &SingleVerifyResp{}
	if err := c.C.ReqAndResp("GET", tUrl, svResp, &opt.ReqOpts); err != nil {
		return nil, err
	}

	return svResp, nil
}
