package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch [URL]",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		repo := args[0]

		fmt.Println("fetch called")

		folder := repo[strings.LastIndex(repo, "/")+1:]

		_, err := git.PlainClone(os.Getenv("HELM_DATA_HOME")+"/starters/"+folder, false, &git.CloneOptions{
			URL:      args[0],
			Progress: os.Stdout,
		})
		if err == nil {
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
