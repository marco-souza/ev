package cmd

import (
	"fmt"
	"strings"

	"github.com/marco-souza/ev/internal/db"
	"github.com/marco-souza/ev/internal/vault"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run  -- [your shell script]",
	Short: "Run a shell script with variables and secrets",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		shellCommand := strings.Join(args, " ")

		v := vault.NewVault()
		conn, err := db.NewConnection(v.DatabaseURI())
		if err != nil {
			panic(err)
		}

		defer conn.Close()

		output, err := v.Run(conn, shellCommand)
		if err != nil {
			panic(err)
		}

		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
