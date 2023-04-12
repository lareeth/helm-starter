package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed starter templates",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir(os.Getenv("HELM_DATA_HOME") + "/starters")
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fmt.Println(file.Name())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
