package workwx

import (
	"log"
	"os"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type defaultLogger struct {
	stderr *log.Logger
	stdout *log.Logger
}

func (l *defaultLogger) Info(msg string) {
	l.stdout.Println(msg)
}

func (l *defaultLogger) Error(msg string) {
	l.stderr.Println(msg)
}

func newDefaultLogger() *defaultLogger {
	stderr := log.New(os.Stderr, "[workwx INFO]", 0)
	stdout := log.New(os.Stdout, "[workwx ERR]", 0)
	return &defaultLogger{stderr: stderr, stdout: stdout}
}
