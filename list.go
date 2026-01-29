package main

import (
	"fmt"
	"log/slog"

	"github.com/alexxyjiang/domain-list-custom/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringP("config", "c", "config.json", "URI of the JSON format config file")
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List available domain lists",
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

		// Process only input to get the list
		for idx, inputConfig := range instance.Config.Input {
			slog.Debug("processing input ...", "processed", idx+1, "total", len(instance.Config.Input), "type", inputConfig.Type, "action", inputConfig.Action)

			converter, err := inputConfig.GetInputConverter()
			if err != nil {
				slog.Error("failed to get input converter", "type", inputConfig.Type, "err", err)
			}

			newContainer, err := converter.Input(instance.Container)
			if err != nil {
				slog.Error("failed to create input container", "type", inputConfig.Type, "err", err)
			}

			if newContainer != nil {
				instance.Container = newContainer
			}
		}
		slog.Debug("all input processors done")

		// List all entries
		fmt.Println("Available domain lists:", instance.Container.Len(), "in total")
		fmt.Println("---")

		names := instance.Container.GetNames()
		for _, name := range names {
			entry, found := instance.Container.GetEntry(name)
			if !found {
				continue
			}
			fmt.Println(" - ", name, "(", len(entry.GetDomains()), "domains)")
		}
	},
}
