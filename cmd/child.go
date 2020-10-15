package cmd

import (
	"github.com/l-lin/go-boilerplate/conf"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewChildCmd() *cobra.Command {
	childCmd := &cobra.Command{
		Use:   "child",
		Short: "Child command description",
		Run:   runChildCmd,
	}
	initChildCmd(childCmd)
	return childCmd
}

func initChildCmd(cmd *cobra.Command) {
	cmd.Flags().String("child-name", "child name", "child name")
}

func runChildCmd(cmd *cobra.Command, args []string) {
	log.Info().Msg("Hello, from child command")
	log.Info().Str("SomeProperty", conf.Get().SomeProperty).Msg("reading config")
	childName, err := cmd.Flags().GetString("child-name")
	if err != nil {
		log.Fatal().Err(err).Msg("could not read flag")
	}
	log.Info().Str("child-name", childName).Msg("reading flag")
}
