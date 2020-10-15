package main

import (
	"github.com/l-lin/go-boilerplate/cmd"
	"github.com/l-lin/go-boilerplate/conf"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{TimeFormat: time.RFC3339, Out: os.Stderr})

	rootCmd := cmd.NewRootCmd(version, buildDate)
	rootCmd.AddCommand(cmd.NewChildCmd())

	command := findCommand(rootCmd)

	cobra.OnInitialize(conf.Init)

	if err := command.Execute(); err != nil {
		log.Err(err).Msg("error when executing the root command")
		os.Exit(1)
	}
}

func findCommand(rootCmd *cobra.Command) *cobra.Command {
	command, _, err := rootCmd.Find(os.Args[1:])
	if err != nil {
		log.Err(err).Msg("could not find command")
		os.Exit(1)
	}
	return command
}
