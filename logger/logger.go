package logger

import (
	"github.com/rs/zerolog"
)

type Logger interface {
	Info() *zerolog.Event
	Debug() *zerolog.Event
	Error() *zerolog.Event
	Err(err error) *zerolog.Event
	Fatal() *zerolog.Event
	Log() *zerolog.Event
	Panic() *zerolog.Event
	Trace() *zerolog.Event
	Warn() *zerolog.Event
	Level() zerolog.Level
	SetLevel(level zerolog.Level)
	SetGlobalLevel(level zerolog.Level)
}

type ZerologLogger struct {
	logger zerolog.Logger
	level  zerolog.Level
}

func NewZerologLogger(logger zerolog.Logger) *ZerologLogger {
	return &ZerologLogger{
		logger: logger,
		level:  zerolog.InfoLevel,
	}
}

func (z *ZerologLogger) Info() *zerolog.Event {
	return z.logger.Info()
}

func (z *ZerologLogger) Debug() *zerolog.Event {
	return z.logger.Debug()
}

func (z *ZerologLogger) Error() *zerolog.Event {
	return z.logger.Error()
}

func (z *ZerologLogger) Err(err error) *zerolog.Event {
	return z.logger.Err(err)
}

func (z *ZerologLogger) Fatal() *zerolog.Event {
	return z.logger.Fatal()
}

func (z *ZerologLogger) Log() *zerolog.Event {
	return z.logger.Log()
}

func (z *ZerologLogger) Panic() *zerolog.Event {
	return z.logger.Panic()
}

func (z *ZerologLogger) Trace() *zerolog.Event {
	return z.logger.Trace()
}

func (z *ZerologLogger) Warn() *zerolog.Event {
	return z.logger.Warn()
}

func (z *ZerologLogger) Level() zerolog.Level {
	return z.level
}

func (z *ZerologLogger) SetLevel(level zerolog.Level) {
	z.logger = z.logger.Level(level)
}

func (z *ZerologLogger) SetGlobalLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
	z.level = level
}
