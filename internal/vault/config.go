package vault

import (
	"errors"
	"os"
	"strings"
)

func (v *Vault) Filepath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	ps := strings.Split(cwd, "/")
	for len(ps) > 0 {
		evHome := append(ps, v.Home)
		fsPath := strings.Join(evHome, "/")

		if ok, _ := dirExists(fsPath); ok {
			return fsPath, nil
		}

		// go up on the fs tree
		ps = ps[:len(ps)-1]
	}

	return "", ErrVaultNotFound
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
