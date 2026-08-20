package cmd

import (
	"context"
	"fmt"

	"github.com/marco-souza/ev/internal/db"
	"github.com/marco-souza/ev/internal/repository"
	"github.com/marco-souza/ev/internal/vault"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [name] [value]",
	Short: "Set a variable name and value",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		value := args[1]

		isSecret, err := cmd.Flags().GetBool("secret")
		if err != nil {
			panic(err)
		}

		v := vault.NewVault()

		ctx := context.Background()
		conn, err := db.NewConnection(v.DatabaseURI())
		if err != nil {
			panic(err)
		}

		defer conn.Close()

		valueType := "variable"
		if isSecret {
			valueType = "secret"
			repo := repository.NewSecretRepository(conn)
			if _, err := repo.Insert(ctx, name, value); err != nil {
				panic(err)
			}
		} else {
			repo := repository.NewVariableRepository(conn)
			if _, err := repo.Insert(ctx, name, value); err != nil {
				panic(err)
			}
		}

		fmt.Printf("%s saved with name=%s\n", valueType, name)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().BoolP("secret", "s", false, "set secret")
}
