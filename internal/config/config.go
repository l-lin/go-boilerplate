package config

import "os"

const defaultUrl = "https://httpbin.org"

// Config of the application
type Config struct {
	// TODO: add the properties needed for your app
	SomeProperty string `yaml:"some_property"`
	Email        string `yaml:"email"`
	URL          string `yaml:"url"`
}

func New() *Config {
	return &Config{
		URL: defaultUrl,
	}
}

func Init(r Repository) error {
	confCreator := &creator{
		Config: New(),
		stdin:  os.Stdin,
		stdout: os.Stdout,
	}
	c, err := confCreator.
		askSomeProperty().
		askEmail().
		create()
	if err != nil {
		return err
	}
	if err = r.Save(c); err != nil {
		return err
	}
	return nil
}
