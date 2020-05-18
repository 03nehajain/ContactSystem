package auth

import (
	"bitbucket.org/exotel/exotel_code/contactsystem/internal/checks"
	"bitbucket.org/exotel/exotel_code/contactsystem/pkg/contracts"
	"bitbucket.org/exotel/exotel_code/contactsystem/pkg/types"
)

//Authenticater defines interface for auth
type Authenticater interface {
	Authenticate(types.Context) (bool, *contracts.Error)
	Authorize(types.Context, ...string) (bool, *contracts.Error)
	GetUserInfo(types.Context) checks.UserInfo
}
