package conf

import (
	"github.com/mitchellh/go-homedir"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

const configName = ".go-boilerplate"

var File string

// Conf of the application
type Conf struct {
	// TODO: add the properties needed for your app
	SomeProperty string `json:"some_property"`
}

// Init the configuration with the given flags
func Init() {
	if File != "" {
		viper.SetConfigFile(File)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal().Err(err).Msg("could not find home directory")
			os.Exit(1)
		}
		viper.SetConfigName(configName)
		viper.SetConfigType("yml")
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Info().Str("cfgFile", viper.ConfigFileUsed()).Msg("using config file")
	}
}

// Get the configuration
func Get() Conf {
	var conf Conf
	if err := viper.Unmarshal(&conf, func(config *mapstructure.DecoderConfig) {
		config.TagName = "json"
		config.WeaklyTypedInput = true
	}); err != nil {
		log.Fatal().Err(err).Msg("could not unmarshal config")
		os.Exit(1)
	}
	return conf
}
