package conf

import "os"

// Conf of the application
type Conf struct {
	// TODO: add the properties needed for your app
	SomeProperty string `yaml:"some_property"`
	Email        string `yaml:"email"`
	URL          string `yaml:"url"`
}

func NewConf() *Conf {
	return &Conf{
		URL: "https://httpbin.org",
	}
}

func (c *Conf) areMandatoryFieldsFilled() bool {
	// TODO: set your mandatory fields
	return c.SomeProperty != "" && c.Email != ""
}

type Provider interface {
	Init() error
	Get() *Conf
}

func NewProvider(r Repository) Provider {
	return &provider{r: r}
}

type provider struct {
	r Repository
}

func (p *provider) Init() error {
	if err := p.r.Init(); err != nil {
		return p.createIfNotExists()
	}
	if err := p.createIfNotExists(); err != nil {
		return err
	}
	return nil
}

func (p *provider) Get() *Conf {
	return p.r.Get()
}

func (p *provider) createIfNotExists() error {
	c := p.Get()
	if c.areMandatoryFieldsFilled() {
		return nil
	}
	confCreator := &creator{
		Conf:   NewConf(),
		stdin:  os.Stdin,
		stdout: os.Stdout,
	}
	var err error
	c, err = confCreator.
		askSomeProperty().
		askEmail().
		create()
	if err != nil {
		return err
	}
	if err = p.r.Save(c); err != nil {
		return err
	}
	return nil
}
