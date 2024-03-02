package controllers

import (
	"fmt"
	"github.com/beebeewijaya-tech/go-todo/internal/entities"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserControllerInterface interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type UserController struct {
	userUsecase entities.UserUsecase
}

func NewUserController(userUsecase entities.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (uc *UserController) Register(c echo.Context) error {
	var user entities.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("error when binding request %v", err))
	}

	u, err := uc.userUsecase.Create(c.Request().Context(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *UserController) Login(c echo.Context) error {
	var user entities.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("error when binding request %v", err))
	}

	u, err := uc.userUsecase.Verify(c.Request().Context(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, u)
}
