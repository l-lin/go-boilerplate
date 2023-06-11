package main

import (
	"fmt"
	"go-boilerplate/internal/config"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewConfigureCmd() *cobra.Command {
	return &cobra.Command{
		Use: "configure",
		Short: fmt.Sprintf(
			"Configure %s CLI options. You will be prompt for configuration values",
			appName,
		),
		Run: runConfigureCmd,
	}
}

func runConfigureCmd(c *cobra.Command, _ []string) {
	if err := config.Init(configRepository); err != nil {
		log.Fatal().
			Stack().
			Err(err).
			Msg("could not init config file")
	}
}
