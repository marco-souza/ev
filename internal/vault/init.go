package vault

import (
	"context"
	"fmt"
	"os"

	"github.com/marco-souza/ev/internal/db"
	"github.com/marco-souza/ev/internal/repository"
)

func (v *Vault) InitDb() error {
	home, err := v.Filepath()
	if err != nil {
		fmt.Println("creating vault")

		if err := os.Mkdir(v.Home, 0o755); err != nil {
			return err
		}

		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		home = fmt.Sprintf("%s/%s", cwd, v.Home)
	}

	filepath := fmt.Sprintf("%s/%s", home, v.Filename)

	fmt.Println("connecting to vault at", filepath)

	conn, err := db.NewConnection(filepath)
	if err != nil {
		return err
	}

	defer conn.Close()

	fmt.Println("creating tables")

	ctx := context.Background()

	variablesRepo := repository.NewVariableRepository(conn)
	if err := variablesRepo.CreateTable(ctx); err != nil {
		return err
	}

	secretsRepo := repository.NewSecretRepository(conn)
	if err := secretsRepo.CreateTable(ctx); err != nil {
		return err
	}

	// rows, err := db.Query(".tables")
	rows, err := conn.Query(`SELECT name FROM sqlite_master
		WHERE type = 'table'
			AND name NOT LIKE 'sqlite_%'
		ORDER BY name
	`)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var table string

		rows.Scan(&table)

		fmt.Printf("  - table: %s\n", table)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}
