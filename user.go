package ne

import "time"

// User represents a ne user.
type User struct {
	ID       int64
	Name     string
	Pass     string
	JoinDate *time.Time
	Admin    string
	Email    string
	Class    string
}

type UserService interface {
	GetAll(limit, offset int) ([]*User, error)
}
