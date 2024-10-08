package auth

import (
	"context"

	"github.com/faridyusof727/e-invoice-go-sdk/configs"
)

type Authenticator interface {
	// LoginAsTaxPayer logs in as a tax payer.
	//
	// It takes a context.Context as a parameter and returns a pointer to a Response and an error.
	LoginAsTaxPayer(ctx context.Context) (*Response, error)

	// LoginAsIntermediary logs in as an intermediary.
	//
	// It takes a context.Context as a parameter.
	// Returns a pointer to a Response and an error.
	LoginAsIntermediary(ctx context.Context, onBehalfOf string) (*Response, error)

	// AccessToken returns the access token for the authenticator.
	//
	// It takes no parameters.
	// Returns a string representing the access token.
	AccessToken() string

	// Config returns the configuration for the authenticator.
	//
	// It takes no parameters.
	// Returns a pointer to a configs.Config.
	Config() *configs.Config
}
