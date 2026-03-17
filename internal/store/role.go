package store

import (
	"context"
	"database/sql"
	"errors"
)

type Role struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       int64  `json:"level"`
}

type RoleStore struct {
	db *sql.DB
}

func (s *RoleStore) GetByName(ctx context.Context, name string) (*Role, error) {
	query := `
	SELECT id, name, level FROM roles WHERE name = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var role Role
	err := s.db.QueryRowContext(ctx, query, name).Scan(&role.ID, &role.Name, &role.Level)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &role, nil
}

func (s *RoleStore) GetByID(ctx context.Context, id int64) (*Role, error) {
	query := `
	SELECT id, name, level FROM roles WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var role Role
	err := s.db.QueryRowContext(ctx, query, id).Scan(&role.ID, &role.Name, &role.Level)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &role, nil
}
