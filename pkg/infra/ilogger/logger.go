package ilogger

type LoggerI interface {
	Debug(msg string, keyvals ...interface{}) error
	Info(msg string, keyvals ...interface{}) error
	Error(msg string, keyvals ...interface{}) error

	With(keyvals ...interface{}) LoggerI
}
