package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Level         string `envconfig:"LOGGER_LEVEL"`
	PrettyConsole bool   `envconfig:"LOGGER_PRETTY_CONSOLE" default:"false"`
}

func Init(c Config, app_name string, app_version string) {
	zerolog.TimeFieldFormat = time.RFC3339

	level, err := zerolog.ParseLevel(c.Level)
	if err != nil {
		level = zerolog.InfoLevel
		log.Warn().Err(err).Msg("Failed to parse log level, using default InfoLevel")
	}
	zerolog.SetGlobalLevel(level)

	log.Logger = log.With().
		Caller().
		Str("app_name", app_name).
		Str("app_version", app_version).
		Logger()

	if c.PrettyConsole {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
	}

	log.Info().Msg("Logger initialized")
}
