package main

import (
	"github.com/l-lin/go-boilerplate/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{TimeFormat: time.RFC3339, Out: os.Stderr})
	cmd.Execute(version, buildDate)
}
