package db

import (
	"context"
	"database/sql"
)

type SecretRepository struct {
	db *sql.DB
}

func NewSecretRepository(db *sql.DB) *SecretRepository {
	return &SecretRepository{db: db}
}

type Secret struct {
	Name           string
	EncryptedValue []byte
}

func (r *SecretRepository) GetByName(ctx context.Context, name string) (*Secret, error) {
	query := "SELECT name, value FROM variables WHERE name = $1"
	row := r.db.QueryRowContext(ctx, query, name)

	var variable Secret
	if err := row.Scan(&variable.Name, &variable.EncryptedValue); err != nil {
		return nil, err
	}

	return &variable, nil
}
