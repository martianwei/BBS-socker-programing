package main

import (
	"bbs/cmd/service"
	"bbs/internal/configs"
	"bbs/internal/logging"
	"log/slog"

	"github.com/spf13/cobra"
)

func main() {
	// database.ConnectDatabase()
	logger := logging.CreateLogger(configs.Cfg.LogLevel)
	slog.SetDefault(logger)

	rootCmd := &cobra.Command{
		Use:   "bbs",
		Short: "bbs is a Bulletin Board System implemented by Go",
		Run: func(cmd *cobra.Command, args []string) {
			slog.Debug("Welcome to bbs !")
		},
	}

	apiserverCmd := &cobra.Command{
		Use:   "apiserver",
		Short: "Run apiserver",
		Run: func(cmd *cobra.Command, args []string) {
			slog.Debug("Hello from apiserver!")
			service.APIServer()
		},
	}

	rootCmd.AddCommand(apiserverCmd)

	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
	}
}
