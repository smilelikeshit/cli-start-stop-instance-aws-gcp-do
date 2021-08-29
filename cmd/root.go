package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{Use: "ectl"}

func Run() error {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	return rootCmd.Execute()
}
