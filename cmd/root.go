package cmd

import (
	"encoding/json"
	"github.com/l-lin/go-boilerplate/ioc"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var container *ioc.Container

// NewRootCmd creates a root command that represents the base command when called without any subcommands
func NewRootCmd(version, buildDate string) *cobra.Command {
	c := &cobra.Command{
		Use:   "go-boilerplate",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run:              runRootCmd,
		PersistentPreRun: initializeIocContainer,
	}
	initRootCmd(c, version, buildDate)
	return c
}

func initRootCmd(c *cobra.Command, version, buildDate string) *pflag.FlagSet {
	c.Version = func(version, buildDate string) string {
		res, err := json.Marshal(map[string]string{"version": version, "build_date": buildDate})
		if err != nil {
			log.Fatal().Err(err).Msg("could not marshal version json")
		}
		return string(res)
	}(version, buildDate)

	c.SetVersionTemplate(`{{printf "%s" .Version}}`)
	c.PersistentFlags().String("config", "", "config file (default will look at $PWD/.go-boilerplate.yml then at $HOME/.go-boilerplate.yml)")
	c.PersistentFlags().Bool("debug", false, "debug mode")

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
	log.Info().Msg("Hello, world")
	log.Info().Str("SomeProperty", container.Conf.SomeProperty).Msg("reading config")
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
}

func initializeIocContainer(c *cobra.Command, _ []string) {
	fp := flagParser{Command: c}
	inputs := ioc.Inputs{
		Viper:    viper.New(),
		Debug:    fp.GetDebug(),
		File:     fp.GetConfigFile(),
		UserName: fp.GetUserName(),
	}
	container = ioc.Boostrap(inputs)
}
