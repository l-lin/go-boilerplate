package main

import (
	"os"
	"time"

	"github.com/carlmjohnson/versioninfo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const appName = "go-boilerplate"

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{TimeFormat: time.RFC3339, Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	rootCmd := NewRootCmd(versioninfo.Version)
	rootCmd.AddCommand(NewConfigureCmd())
	rootCmd.AddCommand(NewCompletionCmd())

	// TODO: add your custom cmd here
	rootCmd.AddCommand(NewUserCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Err(err).Msg("error when executing the root command")
		os.Exit(1)
	}
}
