# Mslm Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/mslmdev/sdk-go)](https://pkg.go.dev/github.com/mslmdev/sdk-go)

The official Go SDK for Mslm APIs.

## Requirements

- Go 1.18 or above.

## Installation

```bash
go get github.com/mslmdev/sdk-go
```

## Setting Up

Start a local Go project if you don't have one:

```bash
go mod init
```

Install the package:

```bash
go get github.com/mslmdev/sdk-go
```

Reference the package's main entrypoint in your program and write some code:

```go
import (
	"fmt"

	"github.com/mslmdev/sdk-go/mslm"
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
