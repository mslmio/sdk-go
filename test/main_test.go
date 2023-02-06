package main_test

import (
	"os"
	"testing"

	"github.com/mslmio/sdk-go/email_verify"
	"github.com/mslmio/sdk-go/mslm"

	"github.com/stretchr/testify/require"
)

var c *mslm.Client

func setup() {
	apiKey := os.Getenv("APIKEY")
	c = mslm.Init(apiKey)
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func Test_EmailVerify_Sv_Real(t *testing.T) {
	resp, err := c.EmailVerify.SingleVerify("support@mslm.io")
	require.NoError(t, err)
	require.Exactly(t, resp.Email, "support@mslm.io")
	require.Exactly(t, resp.Username, "support")
	require.Exactly(t, resp.Domain, "mslm.io")
	require.Exactly(t, resp.Malformed, false)
	require.Exactly(t, resp.Suggestion, "")
	require.Exactly(t, resp.Status, "real")
	require.Exactly(t, resp.HasMailbox, true)
	require.Exactly(t, resp.AcceptAll, false)
	require.Exactly(t, resp.Disposable, false)
	require.Exactly(t, resp.Free, false)
	require.Exactly(t, resp.Role, true)
	require.Exactly(t, resp.Mx, []*email_verify.SingleVerifyRespMx{
		{"ASPMX.L.GOOGLE.COM.", 1},
		{"ALT1.ASPMX.L.GOOGLE.COM.", 5},
		{"ALT2.ASPMX.L.GOOGLE.COM.", 5},
		{"ALT3.ASPMX.L.GOOGLE.COM.", 10},
		{"ALT4.ASPMX.L.GOOGLE.COM.", 10},
	})
}

func Test_EmailVerify_Sv_Fake(t *testing.T) {
	resp, err := c.EmailVerify.SingleVerify("fakefakefake@mslm.io")
	require.NoError(t, err)
	require.Exactly(t, resp.Email, "fakefakefake@mslm.io")
	require.Exactly(t, resp.Username, "fakefakefake")
	require.Exactly(t, resp.Domain, "mslm.io")
	require.Exactly(t, resp.Malformed, false)
	require.Exactly(t, resp.Suggestion, "")
	require.Exactly(t, resp.Status, "fake")
	require.Exactly(t, resp.HasMailbox, false)
	require.Exactly(t, resp.AcceptAll, false)
	require.Exactly(t, resp.Disposable, false)
	require.Exactly(t, resp.Free, false)
	require.Exactly(t, resp.Role, false)
	require.Exactly(t, resp.Mx, []*email_verify.SingleVerifyRespMx{
		{"ASPMX.L.GOOGLE.COM.", 1},
		{"ALT1.ASPMX.L.GOOGLE.COM.", 5},
		{"ALT2.ASPMX.L.GOOGLE.COM.", 5},
		{"ALT3.ASPMX.L.GOOGLE.COM.", 10},
		{"ALT4.ASPMX.L.GOOGLE.COM.", 10},
	})
}

func Test_EmailVerify_Sv_Disposable(t *testing.T) {
	resp, err := c.EmailVerify.SingleVerify("asdfasdf@temp-mail.org")
	require.NoError(t, err)
	require.Exactly(t, resp.Email, "asdfasdf@temp-mail.org")
	require.Exactly(t, resp.Username, "asdfasdf")
	require.Exactly(t, resp.Domain, "temp-mail.org")
	require.Exactly(t, resp.Malformed, false)
	require.Exactly(t, resp.Suggestion, "")
	require.Exactly(t, resp.Status, "real")
	require.Exactly(t, resp.HasMailbox, true)
	require.Exactly(t, resp.AcceptAll, true)
	require.Exactly(t, resp.Disposable, true)
	require.Exactly(t, resp.Free, true)
	require.Exactly(t, resp.Role, false)
	require.Exactly(t, resp.Mx, []*email_verify.SingleVerifyRespMx{
		{"mx.yandex.net.", 10},
	})
}

func Test_EmailVerify_Sv_Malformed(t *testing.T) {
	resp, err := c.EmailVerify.SingleVerify("malformedemail")
	require.NoError(t, err)
	require.Exactly(t, resp.Email, "malformedemail")
	require.Exactly(t, resp.Username, "")
	require.Exactly(t, resp.Domain, "")
	require.Exactly(t, resp.Malformed, true)
	require.Exactly(t, resp.Suggestion, "")
	require.Exactly(t, resp.Status, "fake")
	require.Exactly(t, resp.HasMailbox, false)
	require.Exactly(t, resp.AcceptAll, false)
	require.Exactly(t, resp.Disposable, false)
	require.Exactly(t, resp.Free, false)
	require.Exactly(t, resp.Role, false)
	require.Exactly(t, resp.Mx, []*email_verify.SingleVerifyRespMx{})
}
