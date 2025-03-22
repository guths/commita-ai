package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/guths/commita-ai/core/service"
	"github.com/guths/commita-ai/core/usecase"
	"github.com/guths/commita-ai/internal/adapter"
	"github.com/guths/commita-ai/internal/config"
	"github.com/spf13/cobra"
)

var commitUseCase *usecase.Summarize
var git *service.Git
var commitType service.CommitType

func GetCommitMessage(commitType service.CommitType, summary string) string {
	switch commitType {
	case service.CommitTypeFeat:
		return "✨ feat:\n\n" + summary
	case service.CommitTypeFix:
		return "🐛 fix:\n\n" + summary
	case service.CommitTypeChore:
		return "🛠️ chore:\n\n" + summary
	case service.CommitTypeDocs:
		return "📚 docs:\n\n" + summary
	case service.CommitTypeTest:
		return "✅ test:\n\n" + summary
	default:
		return "🔧 other:\n\n" + summary
	}
}

var cliCmd = &cobra.Command{
	Use:   "c",
	Short: "commit",
	Long:  `Commit`,
	Run: func(cmd *cobra.Command, args []string) {
		if !service.IsValidCommitType(string(commitType)) {
			panic("Invalid commit type")
		}
		diff, err := git.Diff()

		if err != nil {
			panic(err)
		}

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Start()
		s.Suffix = " Summarizing changes... \n"

		summary, err := commitUseCase.Create(commitType, diff)

		s.Stop()

		if err != nil {
			panic(err)
		}

		message := GetCommitMessage(commitType, summary)

		err = git.Commit(message)

		if err != nil {
			fmt.Println("ERRRROOOO", err)
			panic(err)
		}
	},
}

func init() {
	config.LoadEnv()
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP((*string)(&commitType), "type", "t", string(service.CommitTypeFeat), "commit conventional")
	git, _ = service.NewGit()
	api := adapter.NewOpenAiClient()
	ctx := context.Background()
	commitUseCase = usecase.NewSummarize(ctx, api)
}
