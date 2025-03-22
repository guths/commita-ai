package cmd

import (
	"context"

	"github.com/guths/commita-ai/core/service"
	"github.com/guths/commita-ai/core/usecase"
	"github.com/guths/commita-ai/internal/adapter"
	"github.com/guths/commita-ai/internal/config"
	"github.com/spf13/cobra"
)

var commitUseCase *usecase.Summarize
var git *service.Git

var cliCmd = &cobra.Command{
	Use:   "c",
	Short: "commit",
	Long:  `Commit`,
	Run: func(cmd *cobra.Command, args []string) {
		diff, err := git.Diff()

		if err != nil {
			panic(err)
		}

		commitUseCase.Create(diff)
	},
}

func init() {
	config.LoadEnv()
	rootCmd.AddCommand(cliCmd)
	git, _ = service.NewGit()
	api := adapter.NewOpenAiClient()
	ctx := context.Background()
	commitUseCase = usecase.NewSummarize(ctx, api)
}
