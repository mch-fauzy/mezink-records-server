package logger

import (
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// InitLogger initializes the zerolog logger
func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	// format the output as needed here

	log.Logger = log.Output(output)
	log.Trace().Msg("Zerolog initialized")
}

// ErrorWithStack logs and error and its stack trace with custom formatting.
func ErrorWithStack(err error) {
	log.Error().Msgf("%+v", errors.WithStack(err))
}
