package cmd

import (
	"context"

	"github.com/marco-souza/ev/internal/db"
	"github.com/marco-souza/ev/internal/repository"
	"github.com/marco-souza/ev/internal/vault"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [name]",
	Short: "Remove a variable by name",
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
			if err := repo.DeleteByName(ctx, name); err != nil {
				panic(err)
			}

		case false:
			repo := repository.NewVariableRepository(conn)
			if err := repo.DeleteByName(ctx, name); err != nil {
				panic(err)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	rmCmd.Flags().BoolP("secret", "s", false, "save as secret")
}
