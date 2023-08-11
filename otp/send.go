package otp

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/mslmio/sdk-go/lib"
)

type OtpSendReqOpts struct {
	// Common request options.
	lib.ReqOpts
}

type OtpSendResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type OtpSendReq struct {
	Phone         string `json:"phone"`
	TmplSms       string `json:"tmpl_sms"`
	TokenLen      int    `json:"token_len"`
	ExpireSeconds int    `json:"expire_seconds"`
}

func (c *Client) Send(
	otpSendReq *OtpSendReq,
	opts ...*OtpSendReqOpts,
) (*OtpSendResp, error) {
	fmt.Println("here send func start")
	// Prepare options.
	opt := &OtpSendReqOpts{}
	if len(opts) > 0 {
		opt = opts[len(opts)-1]
	}
	opt.ReqOpts = c.C.PrepareOpts(&opt.ReqOpts)

	// Prepare URL.
	qp := map[string]string{}
	tUrl, _ := url.Parse("/api/v1/send")
	tUrl, err := c.C.PrepareUrl("/api/otp/v1/send", qp, &opt.ReqOpts)
	if err != nil {
		return nil, err
	}

	// Marshal.
	data, err := json.Marshal(otpSendReq)
	if err != nil {
		return nil, err
	}

	// Get data.
	otpSendResp := &OtpSendResp{}
	if err := c.C.ReqAndResp("POST", tUrl, data, otpSendResp, &opt.ReqOpts); err != nil {
		return nil, err
	}

	fmt.Println("here send func end")

	return otpSendResp, nil
}
