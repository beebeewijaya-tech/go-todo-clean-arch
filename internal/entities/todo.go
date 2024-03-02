package entities

import (
	"context"
)

// Todo struct will be the main entities being use by the whole project related to todo
// Such as Usecase and Repository
type Todo struct {
	ID     string `json:"id" param:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId string `json:"user_id"`
}

// TodoUsecase will be the main business logic in our application of todo related
type TodoUsecase interface {
	CreateTodo(ctx context.Context, todo Todo, userId string) (Todo, error)
	GetTodo(ctx context.Context, id string, userId string) (Todo, error)
	ListTodo(ctx context.Context, userId string) ([]Todo, error)
	UpdateTodo(ctx context.Context, todo Todo, userId string) (Todo, error)
	DeleteTodo(ctx context.Context, id string, userId string) error
}

// TodoRepository will communicate with our databases regarding todo table
type TodoRepository interface {
	Save(ctx context.Context, todo Todo, userId string) (Todo, error)
	Get(ctx context.Context, id, userId string) (Todo, error)
	List(ctx context.Context, userId string) ([]Todo, error)
	Update(ctx context.Context, todo Todo, userId string) (Todo, error)
	Delete(ctx context.Context, id, userId string) error
}
