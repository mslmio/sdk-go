package email_verify

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/mslmio/sdk-go/lib"
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
	Email      string                 `json:"email" csv:"email"`
	Username   string                 `json:"username" csv:"username"`
	Domain     string                 `json:"domain" csv:"domain"`
	Malformed  bool                   `json:"malformed" csv:"malformed"`
	Suggestion string                 `json:"suggestion" csv:"suggestion"`
	Status     string                 `json:"status" csv:"status"`
	HasMailbox bool                   `json:"has_mailbox" csv:"has_mailbox"`
	AcceptAll  bool                   `json:"accept_all" csv:"accept_all"`
	Disposable bool                   `json:"disposable" csv:"disposable"`
	Free       bool                   `json:"free" csv:"free"`
	Role       bool                   `json:"role" csv:"role"`
	Mx         SingleVerifyRespMxWrap `json:"mx" csv:"mx"`
	Fake       bool                   `json:"fake" csv:"fake"`
	Spamtrap   bool                   `json:"spamtrap" csv:"spamtrap"`
}

type SingleVerifyRespMx struct {
	Host string `json:"host"`
	Pref int    `json:"pref"`
}

type SingleVerifyRespMxWrap []*SingleVerifyRespMx

func (mx SingleVerifyRespMxWrap) MarshalCSV() ([]byte, error) {
	var csvData strings.Builder

	csvData.WriteString("[")
	for i, item := range mx {
		itemStr := fmt.Sprintf(`{"host":"%s","pref":%d}`, item.Host, item.Pref)
		csvData.WriteString(itemStr)

		if i < len(mx)-1 {
			csvData.WriteString(", ")
		}
	}
	csvData.WriteString("]")

	return []byte(csvData.String()), nil
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
	if err := c.C.ReqAndResp("GET", tUrl, nil, svResp, &opt.ReqOpts); err != nil {
		return nil, err
	}

	return svResp, nil
}
