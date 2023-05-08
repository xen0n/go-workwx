package workwx

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type defaultLogger struct{}

func (l *defaultLogger) Info(msg string) {}

func (l *defaultLogger) Error(msg string) {}

func newDefaultLogger() *defaultLogger {
	return &defaultLogger{}
}
