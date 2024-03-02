package controllers

import (
	"fmt"
	"github.com/beebeewijaya-tech/go-todo/internal/entities"
	"github.com/beebeewijaya-tech/go-todo/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TodoControllerInterface interface {
	ListTodo(c echo.Context) error
	GetTodo(c echo.Context) error
	CreateTodo(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type TodoController struct {
	todoUsecase entities.TodoUsecase
	utils       *utils.Utils
}

func NewTodoController(todoUsecase entities.TodoUsecase, utils *utils.Utils) *TodoController {
	return &TodoController{
		todoUsecase: todoUsecase,
		utils:       utils,
	}
}

func (t *TodoController) ListTodo(c echo.Context) error {
	user := t.utils.GetUser(c)
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("error when get user"))
	}

	td, err := t.todoUsecase.ListTodo(c.Request().Context(), user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("error when listing todo %v", err))
	}

	return c.JSON(http.StatusOK, td)
}

func (t *TodoController) DeleteTodo(c echo.Context) error {
	var todo entities.Todo
	if err := c.Bind(&todo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed to bind todo body %v", err))
	}

	user := t.utils.GetUser(c)
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("error when get user"))
	}

	err := t.todoUsecase.DeleteTodo(c.Request().Context(), todo.ID, user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("error when delete todo %v", err))
	}

	return c.JSON(http.StatusOK, "successfully deleted")
}

func (t *TodoController) UpdateTodo(c echo.Context) error {
	var todo entities.Todo
	if err := c.Bind(&todo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed to bind todo body %v", err))
	}

	user := t.utils.GetUser(c)
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("error when get user"))
	}

	td, err := t.todoUsecase.UpdateTodo(c.Request().Context(), todo, user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("error when update todo %v", err))
	}

	return c.JSON(http.StatusOK, td)
}

func (t *TodoController) GetTodo(c echo.Context) error {
	var todo entities.Todo
	if err := c.Bind(&todo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed to bind todo body %v", err))
	}

	user := t.utils.GetUser(c)
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("error when get user"))
	}

	td, err := t.todoUsecase.GetTodo(c.Request().Context(), todo.ID, user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("error when listing todo %v", err))
	}

	return c.JSON(http.StatusOK, td)
}

func (t *TodoController) CreateTodo(c echo.Context) error {
	var todo entities.Todo
	if err := c.Bind(&todo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed to bind todo body %v", err))
	}

	user := t.utils.GetUser(c)
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("error when get user"))
	}

	td, err := t.todoUsecase.CreateTodo(c.Request().Context(), todo, user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("error when submit todo %v", err))
	}

	return c.JSON(http.StatusOK, td)
}
