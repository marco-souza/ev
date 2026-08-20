package cmd

import (
	"github.com/marco-souza/ev/internal/vault"
	"github.com/spf13/cobra"
)

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop vault",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.NewVault()
		if err := v.Drop(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dropCmd)
}
