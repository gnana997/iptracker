package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command {
		Use: "cobra",
		Short: "IPTracker CLI Tool",
		Long: "IPTracker CLI Tool",
	}
)

func Execute() error {
	return rootCmd.Execute()
}