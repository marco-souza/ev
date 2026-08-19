package db

import (
	"context"
	"database/sql"
	"time"

	// libsql driver
	_ "turso.tech/database/tursogo"
)

func NewConnection(filepath string) (*sql.DB, error) {
	db, err := sql.Open("turso", filepath)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
