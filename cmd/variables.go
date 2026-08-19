package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listVarCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all variables envs on this vault",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list variables called")
	},
}

var setVarCmd = &cobra.Command{
	Use:   "set <var> <value>",
	Short: "set an environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set variables called")
	},
}

var getVarCmd = &cobra.Command{
	Use:   "get <var>",
	Short: "get environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get variables called")
	},
}

var rmVarCmd = &cobra.Command{
	Use:   "rm <var>",
	Short: "remove environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rm variables called")
	},
}

// varCmd represents the variables command
var varCmd = &cobra.Command{
	Use:   "variables",
	Short: "variables command",
}

func init() {
	varCmd.AddCommand(setVarCmd)
	varCmd.AddCommand(getVarCmd)
	varCmd.AddCommand(listVarCmd)
	varCmd.AddCommand(rmVarCmd)

	rootCmd.AddCommand(varCmd)
}
