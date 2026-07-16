package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var (
	verbose bool
	rootCmd = &cobra.Command{
		Use:   "cli",
		Short: "Maker Flow Go CLI scaffold",
		Long:  "Minimal Cobra CLI template. Add subcommands under internal/cmd.",
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose logging")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(runCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-cli 0.1.0")
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Demo long-running command with graceful shutdown",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := newLogger()
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		logger.Info("run started")
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				logger.Info("shutting down", "reason", ctx.Err())
				return nil
			case t := <-ticker.C:
				logger.Info("tick", "at", t.Format(time.RFC3339))
			}
		}
	},
}

func newLogger() *slog.Logger {
	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
}
