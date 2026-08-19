package repository

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

func (r *SecretRepository) CreateTable(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS secrets (
		id INTEGER PRIMARY KEY NOT NULL
			DEFAULT (lower(hex(randomblob(16)))),

		name TEXT NOT NULL,
		value BLOB NOT NULL,  -- ciphertext

		UNIQUE (name)
	);`

	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (r *SecretRepository) GetByName(ctx context.Context, name string) (*Secret, error) {
	query := "SELECT name, value FROM secrets WHERE name = ?;"
	row := r.db.QueryRowContext(ctx, query, name)

	var secret Secret
	if err := row.Scan(&secret.Name, &secret.EncryptedValue); err != nil {
		return nil, err
	}

	return &secret, nil
}

func (r *SecretRepository) Insert(ctx context.Context, name, value string) (*Secret, error) {
	// FIXME: encrypt secret
	encryptedValue := value

	query := "INSERT INTO secrets (name, value) VALUES (?, ?) RETURNING *;"
	row := r.db.QueryRowContext(ctx, query, name, encryptedValue)

	var secret Secret
	if err := row.Scan(&secret.Name, &secret.EncryptedValue); err != nil {
		return nil, err
	}

	return &secret, nil
}

func (r *SecretRepository) ListAll(ctx context.Context) (*[]Secret, error) {
	query := "SELECT name FROM secrets;"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var secrets []Secret

	for rows.Next() {
		var v Secret
		if err := rows.Scan(&v.Name); err != nil {
			return nil, err
		}

		secrets = append(secrets, v)
	}

	return &secrets, nil
}

func (r *SecretRepository) DeleteByName(ctx context.Context, name string) error {
	query := "DELETE FROM secrets WHERE name = ?;"
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
