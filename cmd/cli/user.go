package main

import (
	"go-boilerplate/internal/user"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var username string

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
	c.Flags().StringVarP(&username, "username", "u", "username value", "The username")
}

func runUserCmd(c *cobra.Command, _ []string) {
	log.Info().Msg("Hello, from user command")

	conf, err := configRepository.Get()
	if err != nil {
		log.Fatal().Err(err).Msg("could not get the config")
	}

	var ur user.Repository
	ur = user.NewHttpRepository(*conf)

	log.Info().Str("SomeProperty", conf.SomeProperty).Msg("reading config.SomeProperty")
	log.Info().Str("username", username).Msg("reading flag username")

	user, err := ur.Get("userId")
	if err != nil {
		log.Fatal().Str("username", username).Err(err).Msg("could not get the user")
	}
	log.Info().Str("uuid", user.UUID).Msg("got user")
}
