package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewUserCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "user",
		Short: "User command description",
		Run:   runUserCmd,
	}
	initUserCmd(c)
	return c
}

func initUserCmd(c *cobra.Command) {
	c.Flags().String("user-name", "user name", "user name")
}

func runUserCmd(_ *cobra.Command, _ []string) {
	log.Info().Msg("Hello, from user command")
	log.Info().Str("SomeProperty", container.Conf.SomeProperty).Msg("reading config")

	log.Info().Str("user-name", container.UserName).Msg("reading flag")
	user, err := container.UserRepository.Get("userId")
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Info().Str("uuid", user.UUID).Msg("got user")
}
