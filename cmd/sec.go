package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listSecCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all secrets (value reducted)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list secrets called")
	},
}

var setSecCmd = &cobra.Command{
	Use:   "set <secret> <value>",
	Short: "set secret",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set var called")
	},
}

var rmSecCmd = &cobra.Command{
	Use:   "rm <secret>",
	Short: "remove secret",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rm var called")
	},
}

// secCmd represents the sec command
var secCmd = &cobra.Command{
	Use:   "sec",
	Short: "secrets command",
}

func init() {
	secCmd.AddCommand(listSecCmd)
	secCmd.AddCommand(setSecCmd)
	secCmd.AddCommand(rmSecCmd)

	rootCmd.AddCommand(secCmd)
}
