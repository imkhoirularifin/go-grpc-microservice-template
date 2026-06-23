package logger

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Setup(serviceName, level, goEnv string) {
	var l zerolog.Logger
	if goEnv == "development" {
		l = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
			With().
			Timestamp().
			Str("service", serviceName).
			Logger()
	} else {
		l = zerolog.New(os.Stderr).
			With().
			Timestamp().
			Str("service", serviceName).
			Logger()
	}

	l = l.Level(parseLevel(level))
	log.Logger = l
}

func parseLevel(level string) zerolog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zerolog.DebugLevel
	case "warn", "warning":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	default:
		return zerolog.InfoLevel
	}
}
