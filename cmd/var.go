package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// varCmd represents the var command
var varCmd = &cobra.Command{
	Use:   "var",
	Short: "variables command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("var called")
	},
}

func init() {
	rootCmd.AddCommand(varCmd)
}
