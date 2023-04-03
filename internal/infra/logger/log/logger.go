package log

import (
	"log"

	"github.com/d-kv/backend-travel-app/pkg/infra/ilogger"
)

type Logger struct {
	logger log.Logger
}

var _ ilogger.LoggerI = (*Logger)(nil)

func New(l log.Logger) *Logger {
	return &Logger{
		logger: l,
	}
}

func (l *Logger) Debug(msg string, keyvals ...interface{}) error {
	l.logger.Print(msg, keyvals)
	return nil
}

func (l *Logger) Info(msg string, keyvals ...interface{}) error {
	l.logger.Print(msg, keyvals)
	return nil
}

func (l *Logger) Error(msg string, keyvals ...interface{}) error {
	l.logger.Print(msg, keyvals)
	return nil
}

func (l *Logger) With(keyvals ...interface{}) ilogger.LoggerI {
	panic("not implemented") // TODO: Implement
}
