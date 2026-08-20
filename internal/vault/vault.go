package vault

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/marco-souza/ev/internal/db"
	"github.com/marco-souza/ev/internal/repository"
)

type Vault struct {
	Filename string
	Home     string
}

func NewVault() *Vault {
	return &Vault{
		Filename: "vault.db",
		Home:     ".ev",
	}
}

func (v *Vault) DatabaseURI() string {
	return fmt.Sprintf("%s/%s", v.Home, v.Filename)
}

func (v *Vault) InitDb() error {
	filepath, err := v.ensureHomeExists()
	if err != nil {
		return err
	}

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

func (v *Vault) Drop() error {
	if ok, _ := dirExists(v.Home); ok {
		fmt.Println("dropping vault")

		if err := os.RemoveAll(v.Home); err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (v *Vault) ensureHomeExists() (string, error) {
	filepath := v.DatabaseURI()
	if ok, _ := dirExists(filepath); ok {
		return filepath, nil
	}

	fmt.Println("creating vault")

	if err := os.Mkdir(v.Home, 0o755); err != nil {
		return "", err
	}

	return filepath, nil
}

func dirExists(dir string) (bool, error) {
	info, err := os.Stat(dir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, err
		}

		return false, err
	}

	return info.IsDir(), nil
}
