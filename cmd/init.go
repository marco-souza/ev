package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"

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

			fmt.Println("vault not found, creating")
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

			// TODO: create vault.db (sqlite)
			fmt.Println("TODO: initializing database")
		}

		if info != nil && info.IsDir() {
			panic(fmt.Errorf("'%s/%s' exists as a dir", homePath, vaultFile))
		}

		fmt.Println("vault initialized")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
