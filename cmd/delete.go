package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := os.Getenv("HELM_DATA_HOME") + "/starters/" + args[0]

		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			fmt.Println("Starter Doesn't Exist:", args[0])
		} else {
			err := os.RemoveAll(path)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Deleted Starter:", args[0])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
