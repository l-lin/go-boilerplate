package main

import (
	"fmt"
	"go-boilerplate/internal/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	configFile       string
	configRepository config.Repository
	debug            bool
)

// NewRootCmd creates a root command that represents the base command when called without any subcommands
func NewRootCmd(version string) *cobra.Command {
	c := &cobra.Command{
		Use: appName,
		// TODO: update description
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		PersistentPreRun: initCli,
		Run:              runRootCmd,
	}
	initRootCmd(c, version)
	return c
}

func initRootCmd(c *cobra.Command, version string) *pflag.FlagSet {
	c.Version = fmt.Sprintf("%s %s\n", appName, version)
	c.SetVersionTemplate(`{{printf "%s" .Version}}`)

	c.PersistentFlags().StringVarP(
		&configFile,
		"config",
		"c",
		"",
		fmt.Sprintf("config file (default will look at $HOME/.config/%s/config.yml)", appName),
	)
	c.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug mode")

	// TODO: Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	c.PersistentFlags().StringSlice("types", []string{"foo", "bar"}, "types")
	c.PersistentFlags().String("name", "foobar", "name")

	// TODO: Cobra also supports local flags, which will only run when this action is called directly.
	c.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	return c.Flags()
}

func runRootCmd(c *cobra.Command, _ []string) {
	// TODO: choose what you want to do at root level

	// display help if no command provided
	c.Help()

	// or you could perform actions
	log.Info().Msg("Hello, world from rootCmd")

	types, err := c.Flags().GetStringSlice("types")
	if err != nil {
		log.Fatal().Err(err).Msg("could not read flag")
	}

	log.Info().Strs("types", types).Msg("reading flag")
	toggle, err := c.Flags().GetBool("toggle")
	if err != nil {
		log.Fatal().Err(err).Msg("could not read flag")
	}
	log.Info().Bool("toggle", toggle).Msg("reading flag")

	conf, err := configRepository.Get()
	if err != nil {
		log.Fatal().Err(err).Msg("could not get the config")
	}
	log.Info().Str("SomeProperty", conf.SomeProperty).Msg("reading config.SomeProperty")
}

func initCli(c *cobra.Command, _ []string) {
	v := viper.New()
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Str("configFile", configFile).Msg("using config file")
		v.Debug()
	}
	configRepository = config.NewFileRepository(configFile, v)
}
