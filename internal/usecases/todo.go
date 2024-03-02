package usecases

import (
	"context"
	"fmt"
	"github.com/beebeewijaya-tech/go-todo/internal/entities"
)

type TodoUsecase struct {
	todoRepository entities.TodoRepository
}

func NewTodoUsecase(todoRepository entities.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		todoRepository: todoRepository,
	}
}

func (t *TodoUsecase) CreateTodo(ctx context.Context, todo entities.Todo, userId string) (entities.Todo, error) {
	td, err := t.todoRepository.Save(ctx, todo, userId)
	if err != nil {
		return entities.Todo{}, fmt.Errorf("error when submitting to todo repo %v", err)
	}

	return td, err
}

func (t *TodoUsecase) GetTodo(ctx context.Context, id string, userId string) (entities.Todo, error) {
	td, err := t.todoRepository.Get(ctx, id, userId)
	if err != nil {
		return entities.Todo{}, fmt.Errorf("error when get todo repo: %v", err)
	}

	return td, err
}

func (t *TodoUsecase) ListTodo(ctx context.Context, userId string) ([]entities.Todo, error) {
	td, err := t.todoRepository.List(ctx, userId)
	if err != nil {
		return []entities.Todo{}, fmt.Errorf("error when listing to todo repo: %v", err)
	}

	return td, err
}

func (t *TodoUsecase) UpdateTodo(ctx context.Context, todo entities.Todo, userId string) (entities.Todo, error) {
	td, err := t.todoRepository.Update(ctx, todo, userId)
	if err != nil {
		return entities.Todo{}, fmt.Errorf("error when updating to todo repo %v", err)
	}

	return td, err
}

func (t *TodoUsecase) DeleteTodo(ctx context.Context, id string, userId string) error {
	err := t.todoRepository.Delete(ctx, id, userId)
	if err != nil {
		return fmt.Errorf("error when delete todo repo: %v", err)
	}

	return err
}
