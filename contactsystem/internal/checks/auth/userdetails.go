package auth

type userDetails struct {
	userID     string
	accountSID string
	authToken  string
}

func (u userDetails) UserID() string     { return u.userID }
func (u userDetails) AccountSID() string { return u.accountSID }
func (u userDetails) AuthToken() string  { return u.authToken }
