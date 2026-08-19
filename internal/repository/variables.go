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

func (r *VariableRepository) Insert(ctx context.Context, name, value string) (*Variable, error) {
	query := "INSERT INTO variables (name, value) VALUES (?, ?) RETURNING *;"
	row := r.db.QueryRowContext(ctx, query, name, value)

	var variable Variable
	if err := row.Scan(&variable.Name, &variable.Value); err != nil {
		return nil, err
	}

	return &variable, nil
}

func (r *VariableRepository) ListAll(ctx context.Context) (*[]Variable, error) {
	query := "SELECT name, value FROM variables;"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var variables []Variable

	for rows.Next() {
		var v Variable
		if err := rows.Scan(&v.Name, &v.Value); err != nil {
			return nil, err
		}

		variables = append(variables, v)
	}

	// interaction errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &variables, nil
}

func (r *VariableRepository) DeleteByName(ctx context.Context, name string) error {
	query := "DELETE FROM variables WHERE name = $1;"
	result, err := r.db.ExecContext(ctx, query, name)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFount
	}

	return nil
}
