package config

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"io"
	"regexp"
	"strings"
)

// creator is used as a builder to create a new Conf by asking the user the needed information
type creator struct {
	Err error
	*Config
	stdin  io.ReadCloser
	stdout io.WriteCloser
}

func (creator *creator) askSomeProperty() *creator {
	if creator.Err != nil {
		return creator
	}

	prompt := promptui.Prompt{
		Label:    "Some property",
		Validate: validateNotEmpty,
		Stdin:    creator.stdin,
		Stdout:   creator.stdout,
	}
	creator.SomeProperty, creator.Err = prompt.Run()
	return creator
}

func (creator *creator) askEmail() *creator {
	if creator.Err != nil {
		return creator
	}
	prompt := promptui.Prompt{
		Label:    "Email",
		Validate: validateEmail,
		Stdin:    creator.stdin,
		Stdout:   creator.stdout,
	}
	creator.Email, creator.Err = prompt.Run()
	return creator
}

func (creator creator) create() (*Config, error) {
	return creator.Config, creator.Err
}

func validateNotEmpty(s string) error {
	if strings.Trim(s, " ") == "" || strings.Trim(s, "\t") == "" {
		return errors.New("cannot be empty")
	}
	return nil
}

func validateEmail(s string) error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(s) < 3 && len(s) > 254 || !emailRegex.MatchString(s) {
		return fmt.Errorf(`"%s" is not a valid email`, s)
	}
	return nil
}
