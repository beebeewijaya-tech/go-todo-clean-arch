package postgres

import (
	"context"
	"fmt"
	"github.com/beebeewijaya-tech/go-todo/internal/db"
	"github.com/beebeewijaya-tech/go-todo/internal/entities"
	"github.com/google/uuid"
)

type UserRepository struct {
	DatabaseClient *db.Database
}

func NewUserRepository(dc *db.Database) *UserRepository {
	return &UserRepository{
		DatabaseClient: dc,
	}
}

func (u *UserRepository) Save(ctx context.Context, user entities.User) (entities.User, error) {
	user.ID = uuid.New().String()
	query := "INSERT into users (id, email, password) VALUES ($1, $2, $3)"
	_, err := u.DatabaseClient.Client.ExecContext(ctx, query, user.ID, user.Email, user.Password)
	if err != nil {
		return entities.User{}, fmt.Errorf("inserting into users table error %v", err)
	}

	return user, nil
}

func (u *UserRepository) Get(ctx context.Context, user entities.User) (entities.User, error) {
	var us entities.User
	query := "SELECT id, email, password FROM users WHERE email = $1"
	err := u.DatabaseClient.Client.QueryRowxContext(ctx, query, user.Email).StructScan(&us)
	if err != nil {
		return entities.User{}, fmt.Errorf("error when querying user %v", err)
	}

	return us, nil
}
