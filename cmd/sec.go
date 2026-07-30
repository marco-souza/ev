package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// secCmd represents the sec command
var secCmd = &cobra.Command{
	Use:   "sec",
	Short: "secrets command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sec called")
	},
}

func init() {
	rootCmd.AddCommand(secCmd)
}
