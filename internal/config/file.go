package config

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

const (
	appFolderName = "go-boilerplate"
	configName    = "config"
)

// NewFileRepository creates a new FileRepository and set config file to viper instance
func NewFileRepository(file string, v *viper.Viper) Repository {
	fr := &FileRepository{v: v}
	fr.setConfigFile(file)
	return fr
}

// FileRepository persists the app config on a file system using Viper
type FileRepository struct {
	v *viper.Viper
}

// Get the configuration
func (fr *FileRepository) Get() (*Config, error) {
	if err := fr.v.ReadInConfig(); err != nil {
		return nil, errors.Join(err, errors.New("Please use `configure` command to configure your CLI."))
	}

	var conf *Config
	if err := fr.v.Unmarshal(&conf, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
		config.WeaklyTypedInput = true
	}); err != nil {
		log.Fatal().Err(err).Msg("could not unmarshal config")
	}
	return conf, nil
}

// Save the given app config in filesystem
func (fr *FileRepository) Save(c *Config) error {
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

	parent := filepath.Dir(fr.v.ConfigFileUsed())
	if err = os.MkdirAll(parent, 0770); err != nil {
		return err
	}
	return fr.v.WriteConfig()
}

func (fr *FileRepository) setConfigFile(file string) {
	if file != "" {
		fr.v.SetConfigFile(file)
	} else {
		fr.v.SetConfigFile(fr.buildDefaultConfigFile())
	}
}

func (fr *FileRepository) buildDefaultConfigFile() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal().Err(err).Msg("could not fetch the home directory")
	}
	return fmt.Sprintf("%s/.config/%s/%s.yml", home, appFolderName, configName)
}
