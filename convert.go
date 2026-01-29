package main

import (
	"log/slog"

	"github.com/alexxyjiang/domain-list-custom/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.PersistentFlags().StringP("config", "c", "config.json", "URI of the JSON format config file, support both local file path and remote HTTP(S) URL")
}

var convertCmd = &cobra.Command{
	Use:     "convert",
	Aliases: []string{"conv"},
	Short:   "Convert domain list data from one format to another by using config file",
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config")
		slog.Debug("loading config from", "config", configFile)

		instance, err := lib.NewInstance()
		if err != nil {
			slog.Error("failed to create new instance", "err", err)
		}

		if err := instance.InitConfig(configFile); err != nil {
			slog.Error("failed to initial config", "err", err)
		}

		if err := instance.Run(); err != nil {
			slog.Error("failed to convert", "err", err)
		}
		slog.Info("convert success")
	},
}
