package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var cfgCli = &cobra.Command{
	Use:   "config",
	Short: "config",
	Long:  `Commit`,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("failed to get home directory: %w", err)
			os.Exit(1)
		}

		configDir := filepath.Join(homeDir, ".config", "commitaai")

		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			fmt.Println("failed to create config directory: %w", err)
			os.Exit(1)
		}

		configFile := filepath.Join(configDir, "config.yaml")
		defaultConfig := "open_api_key: YOUR-API-KEY\n"

		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			os.WriteFile(configFile, []byte(defaultConfig), 0644)
		}

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(cfgCli)
}
