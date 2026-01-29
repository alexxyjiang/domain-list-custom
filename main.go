package main

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use:   "domain-list-custom",
	Short: "domain-list-custom is a tool to convert and manage domain lists in various formats",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Set level based on flag
		level := slog.LevelInfo
		if verbose {
			level = slog.LevelDebug
		}
		// Configure and set default logger
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
		slog.SetDefault(logger)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable debug logging")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("exit with error", "err", err)
	}
}
