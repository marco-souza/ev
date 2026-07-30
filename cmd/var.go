package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listVarCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all var envs on this vault",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list var called")
	},
}

var setVarCmd = &cobra.Command{
	Use:   "set <var> <value>",
	Short: "set an environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set var called")
	},
}

var getVarCmd = &cobra.Command{
	Use:   "get <var>",
	Short: "get environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get var called")
	},
}

var rmVarCmd = &cobra.Command{
	Use:   "rm <var>",
	Short: "remove environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rm var called")
	},
}

// varCmd represents the var command
var varCmd = &cobra.Command{
	Use:   "var",
	Short: "variables command",
}

func init() {
	varCmd.AddCommand(setVarCmd)
	varCmd.AddCommand(getVarCmd)
	varCmd.AddCommand(listVarCmd)
	varCmd.AddCommand(rmVarCmd)

	rootCmd.AddCommand(varCmd)
}
