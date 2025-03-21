package cmd

import (
	"fmt"

	"github.com/guths/commita-ai/internal"
	"github.com/spf13/cobra"
)

var git *internal.Git

var cliCmd = &cobra.Command{
	Use:   "c",
	Short: "Commit",
	Long:  `Commit`,
	Run: func(cmd *cobra.Command, args []string) {
		stdout, err := git.Diff()

		if err != nil {
			panic(err)
		}

		fmt.Printf(string(stdout))
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	git = internal.NewGit()
}
