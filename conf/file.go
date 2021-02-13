package conf

import (
	"bytes"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
)

const configName = ".go-boilerplate"

// NewFileRepository creates a new FileRepository
func NewFileRepository(file string, v *viper.Viper) Repository {
	return &FileRepository{
		file: file,
		v:    v,
	}
}

// FileRepository persists the app config on a file system using Viper
type FileRepository struct {
	file string
	v    *viper.Viper
}

// Init the configuration with the given flags
func (fr *FileRepository) Init() error {
	if fr.file != "" {
		fr.v.SetConfigFile(fr.file)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			return err
		}
		fr.v.SetConfigName(configName)
		fr.v.SetConfigType("yml")
		fr.v.AddConfigPath(".")
		fr.v.AddConfigPath(home)
	}
	fr.v.AutomaticEnv()
	if err := fr.v.ReadInConfig(); err != nil {
		return err
	}
	log.Debug().Str("cfgFile", fr.v.ConfigFileUsed()).Msg("using config file")
	return nil
}

// Get the configuration
func (fr *FileRepository) Get() *Conf {
	var conf *Conf
	if err := fr.v.Unmarshal(&conf, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
		config.WeaklyTypedInput = true
	}); err != nil {
		log.Fatal().Err(err).Msg("could not unmarshal config")
		os.Exit(1)
	}
	return conf
}

// Save the given app config in filesystem
func (fr *FileRepository) Save(c *Conf) error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	defaultFilePath := fmt.Sprintf("%s/%s.yml", home, configName)
	if fr.v.ConfigFileUsed() == "" {
		fr.v.SetConfigFile(defaultFilePath)
	}

	bb, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	if err = fr.v.ReadConfig(bytes.NewBuffer(bb)); err != nil {
		return err
	}

	log.Debug().
		Str("configFile", fr.v.ConfigFileUsed()).
		Msg("saving config to file")

	if err = fr.v.WriteConfig(); err != nil {
		return err
	}
	if fr.isDefaultFilePath(defaultFilePath) {
		log.Info().
			Msgf(`do not forget to use "--config %s" next time you are using the CLI`, fr.v.ConfigFileUsed())
	}
	return nil
}

func (fr *FileRepository) isDefaultFilePath(defaultFilePath string) bool {
	return fr.v.ConfigFileUsed() != defaultFilePath && fr.v.ConfigFileUsed() != fmt.Sprintf("%s.yml", configName)
}
