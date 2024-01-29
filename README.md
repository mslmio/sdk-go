# Mslm Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/mslmio/sdk-go)](https://pkg.go.dev/github.com/mslmio/sdk-go)

The official Go SDK for Mslm APIs.

## Requirements

- Go 1.18 or above.

## Authentication

Mslm's APIs require an API key. If you don't have one, please read [Authentication](https://mslm.io/docs/api/authentication) to understand how to get an API key before continuing.

## Installation

```bash
go get github.com/mslmio/sdk-go
```

## Usage

Start a local Go project if you don't have one:

```bash
go mod init
```

Install the package:

```bash
go get github.com/mslmio/sdk-go
```

Reference the package's main entrypoint in your program and write some code:

```go
import (
	"fmt"

	"github.com/mslmio/sdk-go/mslm"
)

func main() {
	c := mslm.Init("YOUR API KEY")
	resp, err := c.EmailVerify.SingleVerify("support@mslm.io")
	fmt.Printf("response: %v", resp)
}
```

## About Mslm

mslm focuses on producing world-class business solutions. It’s the
bread-and-butter of our business to prioritize quality on everything we touch.
Excellence is a core value that defines our culture from top to bottom.

[![image](https://avatars.githubusercontent.com/u/50307970?s=200&v=4)](https://mslm.io/)
