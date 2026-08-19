package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
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

		fmt.Println("set called with", name, value, isSecret)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().BoolP("secret", "s", false, "save as secret")
}
