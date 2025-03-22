package cmd

import (
	"context"
	"fmt"
	"os"
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
var customMessage string

func GetCommitMessage(commitType service.CommitType, summary string) string {
	switch commitType {
	case service.CommitTypeFeat:
		return "✨ feat:\n" + summary
	case service.CommitTypeFix:
		return "🐛 fix:\n" + summary
	case service.CommitTypeChore:
		return "🛠️ chore:\n" + summary
	case service.CommitTypeDocs:
		return "📚 docs:\n" + summary
	case service.CommitTypeTest:
		return "✅ test:\n" + summary
	default:
		return "🔧 other:\n" + summary
	}
}

var cliCmd = &cobra.Command{
	Use:   "c",
	Short: "commit",
	Long:  `Commit`,
	Run: func(cmd *cobra.Command, args []string) {
		if !service.IsValidCommitType(string(commitType)) {
			fmt.Println("Error: Invalid commit type")
			os.Exit(1)
		}

		err := git.DiffTest()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		if customMessage != "" {
			message := GetCommitMessage(commitType, customMessage)
			err = git.Commit(message)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			os.Exit(0)
		}

		diff, err := git.Diff()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Start()
		s.Suffix = " Summarizing changes... \n"

		summary, err := commitUseCase.Create(commitType, diff)
		s.Stop()

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		message := GetCommitMessage(commitType, summary)

		err = git.Commit(message)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	config.LoadEnv()
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP((*string)(&commitType), "type", "t", string(service.CommitTypeFeat), "commit conventional")
	cliCmd.Flags().StringVarP(&customMessage, "message", "m", customMessage, "commit conventional")

	git, _ = service.NewGit()
	api := adapter.NewOpenAiClient()
	ctx := context.Background()
	commitUseCase = usecase.NewSummarize(ctx, api)
}
