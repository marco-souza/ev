package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listSecretsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all secrets (value reducted)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list secrets called")
	},
}

var setSecretCmd = &cobra.Command{
	Use:   "set <secret> <value>",
	Short: "set secret",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set secrets called")
	},
}

var rmSecretCmd = &cobra.Command{
	Use:   "rm <secret>",
	Short: "remove secret",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rm secrets called")
	},
}

// secretsCmd represents the sec command
var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "secrets command",
}

func init() {
	secretsCmd.AddCommand(listSecretsCmd)
	secretsCmd.AddCommand(setSecretCmd)
	secretsCmd.AddCommand(rmSecretCmd)

	rootCmd.AddCommand(secretsCmd)
}
