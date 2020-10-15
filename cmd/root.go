package cmd

import (
	"encoding/json"
	"github.com/l-lin/go-boilerplate/conf"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// NewRootCmd creates a root command that represents the base command when called without any subcommands
func NewRootCmd(version, buildDate string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "go-boilerplate",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: runRootCmd,
	}
	initRootCmd(rootCmd, version, buildDate)
	return rootCmd
}

func initRootCmd(rootCmd *cobra.Command, version, buildDate string) *pflag.FlagSet {
	//cobra.OnInitialize(initConfig)
	rootCmd.Version = func(version, buildDate string) string {
		res, err := json.Marshal(map[string]string{"version": version, "build_date": buildDate})
		if err != nil {
			log.Fatal().Err(err).Msg("could not marshal version json")
		}
		return string(res)
	}(version, buildDate)

	rootCmd.SetVersionTemplate(`{{printf "%s" .Version}}`)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&conf.File, "config", "", "config file (default will look at $PWD/.go-boilerplate.yml then at $HOME/.go-boilerplate.yml)")
	rootCmd.PersistentFlags().StringSlice("types", []string{"foo", "bar"}, "types")

	rootCmd.PersistentFlags().String("name", "foobar", "name")

	// Cobra also supports local flags, which will only run when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	return rootCmd.Flags()
}

func runRootCmd(cmd *cobra.Command, args []string) {
	log.Info().Msg("Hello, world")
	log.Info().Str("SomeProperty", conf.Get().SomeProperty).Msg("reading config")
	types, err := cmd.Flags().GetStringSlice("types")
	if err != nil {
		log.Fatal().Err(err).Msg("could not read flag")
	}
	log.Info().Strs("types", types).Msg("reading flag")
	toggle, err := cmd.Flags().GetBool("toggle")
	if err != nil {
		log.Fatal().Err(err).Msg("could not read flag")
	}
	log.Info().Bool("toggle", toggle).Msg("reading flag")
}
