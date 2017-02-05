package jwt

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type Token struct {
	ID       string
	Issuer   string
	Subject  string
	Duration time.Duration
	CSRF     string
}

func Encode(t *Token, secret []byte) (string, error) {
	type MyCustomClaims struct {
		CSRF string `json:"csrf,omitempty"`
		jwt.StandardClaims
	}

	claims := MyCustomClaims{
		CSRF: t.CSRF,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(t.Duration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    t.Issuer,
			Subject:   t.Subject,
			Id:        t.ID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("sign token failed: %v", err)
	}
	return ss, nil
}

func NewUUID() string {
	return uuid.NewV4().String()
}
