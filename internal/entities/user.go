package entities

import (
	"context"
	"github.com/beebeewijaya-tech/go-todo/internal/models"
)

// User struct will be the main entities being use by the whole project related to user
// Such as Usecase and Repository
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserUsecase will be the main business logic in our application of usr
type UserUsecase interface {
	Create(ctx context.Context, user User) (models.UserCreateResponse, error)
	Verify(ctx context.Context, user User) (models.UserLoginResponse, error)
}

// UserRepository will communicate with our databases
type UserRepository interface {
	Save(ctx context.Context, user User) (User, error)
	Get(ctx context.Context, user User) (User, error)
}
