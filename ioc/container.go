package ioc

import (
	"github.com/l-lin/go-boilerplate/conf"
	"github.com/l-lin/go-boilerplate/user"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Container IoC used to bootstrap the beans
type Container struct {
	Inputs
	*conf.Conf
	UserRepository user.Repository
}

// Bootstrap the beans from the given user inputs
func Boostrap(inputs Inputs) *Container {
	container := &Container{
		Inputs: inputs,
	}
	container.registerConf()

	// TODO: add custom beans here
	container.registerUserRepository()

	container.setLogLevel()
	return container
}

func (c *Container) registerUserRepository() {
	var ur user.Repository
	ur = user.NewHttpRepository(*c.Conf, c.Debug)
	c.UserRepository = ur
}

func (c *Container) registerConf() {
	var cr conf.Repository
	cr = conf.NewFileRepository(c.File, c.Viper)
	var cp conf.Provider
	cp = conf.NewProvider(cr)
	if err := cp.Init(); err != nil {
		log.Fatal().
			Err(err).
			Msg("could not initialize config")
	}
	c.Conf = cp.Get()
}

func (c *Container) setLogLevel() {
	if c.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().
			Str("configFile", c.Viper.ConfigFileUsed()).
			Msg("using config file")
		c.Viper.Debug()
	}
}
