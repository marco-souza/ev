package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Vault struct {
	vaultfile string
}

func NewVault() *Vault {
	return &Vault{
		vaultfile: "vault.db",
	}
}

func (v *Vault) Filepath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ps := strings.Split(cwd, "/")
	for len(ps) > 0 {
		evHome := append(ps, ".ev", v.vaultfile)
		fsPath := strings.Join(evHome, "/")

		if dirExists(fsPath) {
			fmt.Println("vault found at:", fsPath)
			return fsPath, nil
		}

		// go up on the fs tree
		ps = ps[:len(ps)-1]
	}

	fmt.Println("vault not found")
	return "", nil
}

func dirExists(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}

		panic(err)
	}

	return info.IsDir()
}
