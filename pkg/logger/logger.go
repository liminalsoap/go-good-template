package logger

import (
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

type Interface interface {
	Debug(message string)
	Debugf(message string, args ...interface{})
	Info(message string)
	Infof(message string, args ...interface{})
	Warn(message string)
	Warnf(message string, args ...interface{})
	Error(message string)
	Errorf(message string, args ...interface{})
	Fatal(message string)
	Fatalf(message string, args ...interface{})
}

type Logger struct {
	logger *zerolog.Logger
}

func NewLogger(logLevel string) *Logger {
	var l zerolog.Level
	switch strings.ToLower(logLevel) {
	case "debug":
		l = zerolog.DebugLevel
	case "info":
		l = zerolog.InfoLevel
	case "warn":
		l = zerolog.WarnLevel
	case "error":
		l = zerolog.ErrorLevel
	case "fatal":
		l = zerolog.FatalLevel
	}
	zerolog.SetGlobalLevel(l)
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return &Logger{
		logger: &logger,
	}
}

func (l Logger) Debug(message string) {
	l.logger.Debug().Msg(message)
}

func (l Logger) Debugf(message string, args ...interface{}) {
	l.logger.Debug().Msgf(message, args...)
}

func (l Logger) Info(message string) {
	l.logger.Info().Msg(message)
}

func (l Logger) Infof(message string, args ...interface{}) {
	l.logger.Info().Msgf(message, args...)
}

func (l Logger) Warn(message string) {
	l.logger.Warn().Msg(message)
}

func (l Logger) Warnf(message string, args ...interface{}) {
	l.logger.Warn().Msgf(message, args...)
}

func (l Logger) Error(message string) {
	l.logger.Error().Msg(message)
}

func (l Logger) Errorf(message string, args ...interface{}) {
	l.logger.Error().Msgf(message, args...)
}

func (l Logger) Fatal(message string) {
	l.logger.Fatal().Msg(message)
}

func (l Logger) Fatalf(message string, args ...interface{}) {
	l.logger.Fatal().Msgf(message, args...)
}
