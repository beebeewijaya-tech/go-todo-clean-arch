package postgres

import (
	"context"
	"fmt"
	"github.com/beebeewijaya-tech/go-todo/internal/db"
	"github.com/beebeewijaya-tech/go-todo/internal/entities"
	"github.com/google/uuid"
)

type TodoRepository struct {
	DatabaseClient *db.Database
}

func NewTodoRepository(db *db.Database) *TodoRepository {
	return &TodoRepository{
		DatabaseClient: db,
	}
}

func (t *TodoRepository) Save(ctx context.Context, todo entities.Todo, userId string) (entities.Todo, error) {
	todo.ID = uuid.New().String()
	query := "INSERT INTO todos (id, title, body, user_id) VALUES ($1, $2, $3, $4)"
	_, err := t.DatabaseClient.Client.ExecContext(ctx, query, todo.ID, todo.Title, todo.Body, userId)
	if err != nil {
		return entities.Todo{}, fmt.Errorf("error when inserting todo table %v", err)
	}

	return todo, nil
}

func (t *TodoRepository) Get(ctx context.Context, id, userId string) (entities.Todo, error) {
	var td entities.Todo
	query := "SELECT id, title, body FROM todos WHERE id = $1 AND user_id = $2"
	err := t.DatabaseClient.Client.QueryRowxContext(ctx, query, id, userId).StructScan(&td)
	if err != nil {
		return entities.Todo{}, fmt.Errorf("error when selecting todo table %v", err)
	}

	return td, nil
}

func (t *TodoRepository) List(ctx context.Context, userId string) ([]entities.Todo, error) {
	var td []entities.Todo
	query := "SELECT id, title, body FROM todos WHERE user_id = $1"
	rows, err := t.DatabaseClient.Client.QueryxContext(ctx, query, userId)
	if err != nil {
		return []entities.Todo{}, fmt.Errorf("error when selecting todo table %v", err)
	}

	for rows.Next() {
		var tt entities.Todo
		err := rows.StructScan(&tt)
		if err != nil {
			return []entities.Todo{}, fmt.Errorf("error when structing todo table %v", err)
		}

		td = append(td, tt)
	}

	return td, nil
}

func (t *TodoRepository) Update(ctx context.Context, todo entities.Todo, userId string) (entities.Todo, error) {
	query := "UPDATE todos SET title = $2, body = $3 WHERE id = $1 AND user_id = $4"
	_, err := t.DatabaseClient.Client.ExecContext(ctx, query, todo.ID, todo.Title, todo.Body, userId)
	if err != nil {
		return entities.Todo{}, fmt.Errorf("error when updating todo table %v", err)
	}

	return todo, nil
}

func (t *TodoRepository) Delete(ctx context.Context, id string, userId string) error {
	query := "DELETE FROM todos WHERE id = $1 AND user_id = $2"
	_, err := t.DatabaseClient.Client.ExecContext(ctx, query, id, userId)
	if err != nil {
		return fmt.Errorf("error when updating todo table %v", err)
	}

	return nil
}
