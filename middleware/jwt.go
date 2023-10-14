package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

// the 'exp' {expiration time} claims. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4
//ExpiresAt *NumericDate `json:"exp,omitempty"`

// the 'exp' {expiration time} claims. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.5
//NotBefore *NumericDate `json:"nbf,omitempty"`

// the 'exp' {expiration time} claims. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.6
//IssuedAt *NumericDate `json:"iat,omitempty"`

// the 'exp' {expiration time} claims. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.7

func GenerateJWTToken(id int, name string) string {
	claims := &jwtCustomClaims{
		id,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// ini anggap tidak ada error. Sblmnya: "t, err := ....."
	t, _ := token.SignedString([]byte("123"))
	return t
}
