package utils

import (
	"github.com/beebeewijaya-tech/go-todo/internal/entities"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	jwt.RegisteredClaims
}

// GenerateToken will generate a token based on user credentials
func (u *Utils) GenerateToken(user entities.User) string {
	claims := &CustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    user.ID,
			Subject:   user.Email,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedKey := u.config.GetString("jwt.secret")
	ss, err := token.SignedString([]byte(signedKey))
	if err != nil {
		return ""
	}
	return ss
}
