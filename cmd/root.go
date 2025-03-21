package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

var rootCmd = &cobra.Command{
	Use:   "cai",
	Short: "Nothing",
	Long:  `Nothing`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running default action for myapp")
	},
}
