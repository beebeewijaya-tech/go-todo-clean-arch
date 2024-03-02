package utils

import (
	"github.com/beebeewijaya-tech/go-todo/internal/entities"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (u *Utils) GetUser(c echo.Context) *entities.User {
	token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		return &entities.User{}
	}
	claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	if !ok {
		return &entities.User{}
	}

	issuer, err := claims.GetIssuer()
	if err != nil {
		return &entities.User{}
	}

	email, err := claims.GetSubject()
	if err != nil {
		return &entities.User{}
	}

	return &entities.User{
		ID:    issuer,
		Email: email,
	}
}
