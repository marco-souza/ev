package db

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func getVaultPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ps := strings.Split(cwd, "/")
	for len(ps) > 0 {
		evHome := append(ps, ".ev/vault.db")
		fsPath := strings.Join(evHome, "/")
		fmt.Println(fsPath)

		if dirExists(fsPath) {
			return fsPath, nil
		}

		// go up on the fs tree
		ps = ps[:len(ps)-1]
	}

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
