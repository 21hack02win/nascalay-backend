package log

type Logger interface {
	Infof(format string, args ...interface{})
	Error(i ...interface{})
}