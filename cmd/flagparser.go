package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type flagParser struct {
	*cobra.Command
	err error
}

func (fp flagParser) GetString(key string, panicIfError bool) string {
	s, err := fp.Flags().GetString(key)
	if err != nil && panicIfError {
		log.Fatal().Err(err).Msg("could not parse flag")
		panic(err)
	}
	return s
}

func (fp flagParser) GetBool(key string, panicIfError bool) bool {
	b, err := fp.Flags().GetBool(key)
	if err != nil && panicIfError {
		log.Fatal().Err(err).Msg("could not parse flag")
		panic(err)
	}
	return b
}

func (fp flagParser) GetConfigFile() string {
	return fp.GetString("config", true)
}

func (fp flagParser) GetDebug() bool {
	return fp.GetBool("debug", true)
}

func (fp flagParser) GetUserName() string {
	return fp.GetString("user-name", false)
}
