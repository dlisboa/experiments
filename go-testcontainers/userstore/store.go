package userstore

import (
	"database/sql"
	"fmt"
)

type queryable interface {
	Exec(string, ...any) (sql.Result, error)
	QueryRow(string, ...any) *sql.Row
}

type Store struct {
	db queryable
}

type User struct {
	Name string
}

func (s *Store) Insert(name string) error {
	query := `INSERT INTO users (name) VALUES ($1)`

	_, err := s.db.Exec(query, name)
	if err != nil {
		return fmt.Errorf("userstore: insert: %w", err)
	}
	return nil
}

func (s *Store) GetByName(name string) (User, error) {
	query := `SELECT * FROM users WHERE name = $1 LIMIT 1`

	var u User
	row := s.db.QueryRow(query, name)
	err := row.Scan(&u.Name)
	if err != nil {
		return User{}, fmt.Errorf("usersstore: GetByName: %w", err)
	}
	return u, nil
}
