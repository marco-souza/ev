package db

import (
	"database/sql"
	"fmt"

	_ "turso.tech/database/tursogo"
)

var dbPath = ".ev/vault.db"

func getDatabase() *sql.DB {
	conn, err := sql.Open("turso", dbPath)
	if err != nil {
		panic(err)
	}

	return conn
}

func InitDb() {
	db := getDatabase()
	defer db.Close()

	fmt.Println("creating tables")

	db.Exec(`CREATE TABLE IF NOT EXISTS projects (
		id		TEXT PRIMARY KEY NOT NULL DEFAULT (lower(hex(randomblob(16)))),
		name	TEXT NOT NULL
	)`)

	db.Exec(`CREATE TABLE IF NOT EXISTS variables (
		id TEXT PRIMARY KEY NOT NULL
			DEFAULT (lower(hex(randomblob(16)))),

		project_id TEXT NOT NULL
			REFERENCES projects(id) ON DELETE CASCADE,

		key TEXT NOT NULL,
		value TEXT NOT NULL,

		UNIQUE (project_id, key)
	)`)

	db.Exec(`CREATE TABLE IF NOT EXISTS secrets (
		id TEXT PRIMARY KEY NOT NULL
			DEFAULT (lower(hex(randomblob(16)))),

		project_id TEXT NOT NULL
			REFERENCES projects(id) ON DELETE CASCADE,

		key TEXT NOT NULL,
		encrypted_value BLOB NOT NULL, -- ciphertext

		UNIQUE (project_id, key)
	)`)

	// rows, err := db.Query(".tables")
	rows, err := db.Query(`SELECT name FROM sqlite_master
		WHERE type = 'table'
			AND name NOT LIKE 'sqlite_%'
		ORDER BY name
	`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var table string

		rows.Scan(&table)

		fmt.Printf("  - table: %s\n", table)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}
}
