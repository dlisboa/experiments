package users

import (
	"context"
	"fmt"

	"demo/internal/db"
)

type Store struct {
	db db.DBTX
}

type User struct {
	Name string
}

func (s *Store) Insert(ctx context.Context, name string) error {
	query := `INSERT INTO users (name) VALUES ($1)`

	_, err := s.db.ExecContext(ctx, query, name)
	if err != nil {
		return fmt.Errorf("users: store.Insert: %w", err)
	}
	return nil
}

func (s *Store) GetByName(ctx context.Context, name string) (User, error) {
	query := `SELECT * FROM users WHERE name = $1 LIMIT 1`

	var u User
	row := s.db.QueryRowContext(ctx, query, name)
	err := row.Scan(&u.Name)
	if err != nil {
		return User{}, fmt.Errorf("users: store.GetByName: %w", err)
	}
	return u, nil
}
