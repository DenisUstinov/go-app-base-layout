package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Level         string `envconfig:"LOGGER_LEVEL|info" default:"info"`
	PrettyConsole bool   `envconfig:"LOGGER_PRETTY_CONSOLE|false" default:"false"`
}

func Init(c Config, app_name string, app_version string) {
	zerolog.TimeFieldFormat = time.RFC3339

	level, err := zerolog.ParseLevel(c.Level)
	if err != nil {
		log.Fatal().Err(err).Msgf("Invalid log level specified: %s. Exiting.", c.Level)
	}
	zerolog.SetGlobalLevel(level)

	log.Logger = log.With().
		Caller().
		Str("app_name", app_name).
		Str("app_version", app_version).
		Logger()

	if c.PrettyConsole {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out: os.Stderr, TimeFormat: "15:04:05"})
	}

	log.Info().Msg("Logger initialized")
}
