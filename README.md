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

## Additional Resources

See the official [API Reference Documentation](https://mslm.io/docs/api) for
details on each API's actual interface, which is implemented by this SDK.

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for more information.

## Security

See [Security Issue
Notifications](CONTRIBUTING.md#security-issue-notifications) for more
information.

## License

This project is licensed under the [MIT License](LICENSE).

## About Mslm

At Mslm, we're all about making things better. Where others see norm, we see
opportunity.

We build world-class solutions to the toughest problems. Excellence is a core
value that defines our approach from top to bottom.

We're all about creating positive value for ourselves, our partners and our
societies.

We do it by focusing on quality and the long-term, while intelligently
maneuvering the complex realities of day-to-day business.

Our partners share our perspective, and we jointly produce truly world-class
solutions to the toughest problems.

[![image](https://avatars.githubusercontent.com/u/50307970?s=200&v=4)](https://mslm.io/)
