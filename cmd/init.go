package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/marco-souza/ev/internal/db"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a vault",
	Run: func(cmd *cobra.Command, args []string) {
		homePath := ".ev"
		vaultFile := "vault.db"

		// create .ev folder
		info, err := os.Stat(homePath)
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				panic(err)
			}

			fmt.Println("creating vault")
			err := os.Mkdir(homePath, 0o755)
			if err != nil {
				panic(err)
			}
		}

		if info != nil && !info.IsDir() {
			panic(fmt.Errorf("'%s' exists as a file", homePath))
		}

		info, err = os.Stat(path.Join(homePath, vaultFile))
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				panic(err)
			}

			// INFO: create vault.db (sqlite)
			db.InitDb()
		}

		if info != nil && info.IsDir() {
			panic(fmt.Errorf("'%s/%s' exists as a dir", homePath, vaultFile))
		}

		// INFO: append to .gitignore file
		f, err := os.OpenFile(".gitignore", os.O_CREATE|os.O_RDWR, 0o644)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		sc := bufio.NewScanner(f)
		for sc.Scan() {
			line := sc.Text()
			fmt.Println(line)
			if line == homePath {
				// INFO: skip write if already exits
				return
			}
		}

		if err = sc.Err(); err != nil {
			panic(err)
		}

		fmt.Println("appending to .gitignore")
		_, err = f.WriteString(homePath + "\n")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
