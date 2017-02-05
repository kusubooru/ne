package ne

import (
	"time"

	"golang.org/x/net/context"
)

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
	GetAll(ctx context.Context, limit, offset int64) ([]*User, error)
}
