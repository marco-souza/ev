package cmd

import (
	"context"
	"fmt"

	"github.com/marco-souza/ev/internal/db"
	"github.com/marco-souza/ev/internal/repository"
	"github.com/marco-souza/ev/internal/vault"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List vault values",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
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
			secrets, err := repo.ListAll(ctx)
			if err != nil {
				panic(err)
			}

			for _, variable := range secrets {
				fmt.Printf("%s=%s\n", variable.Name, variable.Value)
			}

		case false:
			repo := repository.NewVariableRepository(conn)
			values, err := repo.ListAll(ctx)
			if err != nil {
				panic(err)
			}

			for _, variable := range values {
				fmt.Printf("%s=%s\n", variable.Name, variable.Value)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	lsCmd.Flags().BoolP("secret", "s", false, "list secret")
}
