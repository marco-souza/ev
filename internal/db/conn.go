package db

import (
	"fmt"
	"os"
	"strings"
)

func Path() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ps := strings.Split(cwd, "/")
	for len(ps) > 0 {
		evHome := append(ps, ".ev/vault.db")
		fsPath := strings.Join(evHome, "/")

		if dirExists(fsPath) {
			return fsPath, nil
		}
		fmt.Println(fsPath)

		// go up on the fs tree
		ps = ps[:len(ps)-1]
	}

	return "", nil
}

func dirExists(dir string) bool {
	// TODO: implement dir exists
	return len(dir) == 0
}
