package domain

type User struct {
	ID       int
	Username string
	Password string
}

type UserRepository interface {
	GetByUsername(username string) (*User, error)
	Save(user *User) error
}

type Authenticator interface {
	Authenticate(username, password string) (string, error)
}