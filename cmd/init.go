package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/marco-souza/ev/internal/vault"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a vault",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.NewVault()
		if err := v.InitDb(); err != nil {
			panic(err)
		}

		// INFO: append to .gitignore file
		f, err := os.OpenFile(".gitignore", os.O_CREATE|os.O_RDWR, 0o644)

		if err != nil {
			panic(err)
		}
		defer f.Close()

		sc := bufio.NewScanner(f)
		for sc.Scan() {
			line := sc.Text()
			fmt.Println(line)

			if line == v.Home {
				// INFO: skip write if already exits
				return
			}
		}

		if err = sc.Err(); err != nil {
			panic(err)
		}

		fmt.Println("appending to .gitignore")
		_, err = f.WriteString(v.Home + "\n")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
