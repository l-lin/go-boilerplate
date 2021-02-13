package main

import (
	"github.com/l-lin/go-boilerplate/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{TimeFormat: time.RFC3339, Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	rootCmd := cmd.NewRootCmd(version, buildDate)
	rootCmd.AddCommand(cmd.NewUserCmd())
	rootCmd.AddCommand(cmd.NewCompletionCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Err(err).Msg("error when executing the root command")
		os.Exit(1)
	}
}
