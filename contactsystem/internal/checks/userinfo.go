package checks

// UserInfo defines interface that holds user info
type UserInfo interface {
	UserID() string
	AccountSID() string
}
