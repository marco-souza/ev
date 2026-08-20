package cmd

import (
	"context"
	"fmt"

	"github.com/marco-souza/ev/internal/db"
	"github.com/marco-souza/ev/internal/repository"
	"github.com/marco-souza/ev/internal/vault"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [name]",
	Short: "Get a variable value by name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

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

		switch isSecret {
		case true:
			repo := repository.NewSecretRepository(conn)
			secret, err := repo.GetByName(ctx, name)
			if err != nil {
				panic(err)
			}

			fmt.Println(secret.Value)

		case false:
			repo := repository.NewVariableRepository(conn)
			value, err := repo.GetByName(ctx, name)
			if err != nil {
				panic(err)
			}

			fmt.Println(value.Value)

		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().BoolP("secret", "s", false, "save as secret")
}
