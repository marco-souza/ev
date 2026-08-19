package db

import (
	"context"
	"database/sql"
)

type VariableRepository struct {
	db *sql.DB
}

func NewVariableRepository(db *sql.DB) *VariableRepository {
	return &VariableRepository{db: db}
}

type Variable struct {
	Name  string
	Value string
}

func (r *VariableRepository) GetByName(ctx context.Context, name string) (*Variable, error) {
	query := "SELECT name, value FROM variables WHERE name = $1"
	row := r.db.QueryRowContext(ctx, query, name)

	var variable Variable
	if err := row.Scan(&variable.Name, &variable.Value); err != nil {
		return nil, err
	}

	return &variable, nil
}
