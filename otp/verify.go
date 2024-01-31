package otp

import (
	"encoding/json"

	"github.com/mslmio/sdk-go/lib"
)

type OtpTokenVerifyReqOpts struct {
	// Common request options.
	lib.ReqOpts
}

type OtpTokenVerifyResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type OtpTokenVerifyReq struct {
	Phone   string `json:"phone"`
	Token   string `json:"token"`
	Consume *bool  `json:"consume"`
}

func (c *Client) Verify(
	otpTokenVerifyReq *OtpTokenVerifyReq,
	opts ...*OtpTokenVerifyReqOpts,
) (*OtpTokenVerifyResp, error) {
	// Prepare options.
	opt := &OtpTokenVerifyReqOpts{}
	if len(opts) > 0 {
		opt = opts[len(opts)-1]
	}
	opt.ReqOpts = c.C.PrepareOpts(&opt.ReqOpts)

	// Prepare URL.
	qp := map[string]string{}
	tUrl, err := c.C.PrepareUrl("/api/otp/v1/token_verify", qp, &opt.ReqOpts)
	if err != nil {
		return nil, err
	}

	// Marshal.
	data, err := json.Marshal(otpTokenVerifyReq)
	if err != nil {
		return nil, err
	}

	// Get data.
	otpTokenVerifyResp := &OtpTokenVerifyResp{}
	if err := c.C.ReqAndResp("POST", tUrl, data, otpTokenVerifyResp, &opt.ReqOpts); err != nil {
		return nil, err
	}

	return otpTokenVerifyResp, nil
}
