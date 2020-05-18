package auth

import (
	"bitbucket.org/exotel/exotel_code/contactsystem/internal/checks"

	"bitbucket.org/exotel/exotel_code/contactsystem/pkg/contracts"
	"bitbucket.org/exotel/exotel_code/contactsystem/pkg/types"
)

type basicAuth struct {
	password   string
	userID     string
	accountSID string
}

func (ba basicAuth) UserID() string     { return ba.userID }
func (ba basicAuth) AccountSID() string { return ba.accountSID }

// Authorize checks if scopes are available for the given user
func (ba basicAuth) Authorize(ctx types.Context, scopes ...string) (bool, *contracts.Error) {
	return true, nil
}

// GetUserInfo returns user information
func (ba basicAuth) GetUserInfo(ctx types.Context) checks.UserInfo {
	return ba
}

// BasicAuth authenticator function for http request
func BasicAuth(ctr *contracts.BasicAuthCredentials) Authenticater {
	return basicAuth{userID: *ctr.UserName, password: *ctr.Password}
}

// Authenticate checks if the user exists with the given credentials
func (ba basicAuth) Authenticate(ctx types.Context) (bool, *contracts.Error) {
	// TODO:Update this
	// TODO: Document default username and password : thenga, manga

	return true, nil
}
