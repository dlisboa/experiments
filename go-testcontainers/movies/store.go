package movies

import (
	"context"
	"fmt"

	"demo/internal/db"
)

type Store struct {
	db db.DBTX
}

type Movie struct {
	Name string
}

func (s *Store) Insert(ctx context.Context, name string) error {
	query := `INSERT INTO movies (name) VALUES ($1)`

	_, err := s.db.ExecContext(ctx, query, name)
	if err != nil {
		return fmt.Errorf("movies: store.Insert: %w", err)
	}
	return nil
}

func (s *Store) GetByName(ctx context.Context, name string) (Movie, error) {
	query := `SELECT * FROM movies WHERE name = $1 LIMIT 1`

	var u Movie
	row := s.db.QueryRowContext(ctx, query, name)
	err := row.Scan(&u.Name)
	if err != nil {
		return Movie{}, fmt.Errorf("movies: store.GetByName: %w", err)
	}
	return u, nil
}
