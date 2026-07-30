package cmd

import (
	"fmt"

	"github.com/marco-souza/ev/internal/db"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a vault",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")

		// TODO: create .ev folder
		// TODO: add .ev to .gitignore
		// TODO: create vault.db (sqlite)

		db.Path()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
