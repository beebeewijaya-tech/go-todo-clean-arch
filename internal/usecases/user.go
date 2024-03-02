package usecases

import (
	"context"
	"fmt"
	"github.com/beebeewijaya-tech/go-todo/internal/entities"
	"github.com/beebeewijaya-tech/go-todo/internal/models"
	"github.com/beebeewijaya-tech/go-todo/internal/utils"
)

type UserUsecase struct {
	UserRepository entities.UserRepository
	Utils          *utils.Utils
}

func NewUserUsecase(userRepo entities.UserRepository, utilities *utils.Utils) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepo,
		Utils:          utilities,
	}
}

func (u *UserUsecase) Create(ctx context.Context, user entities.User) (models.UserCreateResponse, error) {
	user.Password = u.Utils.HashPassword(user.Password)
	us, err := u.UserRepository.Save(ctx, user)
	if err != nil {
		return models.UserCreateResponse{}, fmt.Errorf("error when submitting user to the repository: %v", err)
	}

	return models.UserCreateResponse{
		ID:    us.ID,
		Email: us.Email,
	}, nil
}

func (u *UserUsecase) Verify(ctx context.Context, user entities.User) (models.UserLoginResponse, error) {
	us, err := u.UserRepository.Get(ctx, user)
	if err != nil {
		return models.UserLoginResponse{}, fmt.Errorf("error when submitting user to the repository: %v", err)
	}

	err = u.Utils.VerifyPassword(user.Password, us.Password)
	if err != nil {
		return models.UserLoginResponse{}, fmt.Errorf("error when verifying password: %v", err)
	}

	token := u.Utils.GenerateToken(us)

	return models.UserLoginResponse{
		ID:    us.ID,
		Email: us.Email,
		Token: token,
	}, nil
}
