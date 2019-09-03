package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	std *logrus.Logger
	err *logrus.Logger
}

func newFileLogger(filepath, level string) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	log.SetLevel(lvl)

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	return log
}

func New(serviceName, path, perfix, level string) *Logger {
	if perfix == "" {
		stdFilename := fmt.Sprintf("%s.app.log", serviceName)
		std := newFileLogger(filepath.Join(path, stdFilename), level)
		errFilename := fmt.Sprintf("%s.err.log", serviceName)
		err := newFileLogger(filepath.Join(path, errFilename), "error")
		return &Logger{std, err}
	}
	stdFilename := fmt.Sprintf("%s.%s.log", serviceName, perfix)
	std := newFileLogger(filepath.Join(path, stdFilename), level)
	errFilename := fmt.Sprintf("%s.%s.err.log", serviceName, perfix)
	err := newFileLogger(filepath.Join(path, errFilename), "error")
	return &Logger{std, err}
}

func (log *Logger) Debug(args ...interface{}) {
	log.std.Debug(args...)
}

func (log *Logger) Debugf(format string, args ...interface{}) {
	log.std.Debugf(format, args...)
}

func (log *Logger) Info(args ...interface{}) {
	log.std.Info(args...)
}

func (log *Logger) Infof(format string, args ...interface{}) {
	log.std.Infof(format, args...)
}

func (log *Logger) Warn(args ...interface{}) {
	log.std.Warn(args...)
}

func (log *Logger) Warnf(format string, args ...interface{}) {
	log.std.Warnf(format, args...)
}

func (log *Logger) Error(args ...interface{}) {
	log.std.Error(args...)
	log.err.Error(args...)
}

func (log *Logger) Errorf(format string, args ...interface{}) {
	log.std.Errorf(format, args...)
	log.err.Errorf(format, args...)
}

func (log *Logger) Panic(args ...interface{}) {
	log.std.Panic(args...)
	log.err.Panic(args...)
}

func (log *Logger) Panicf(format string, args ...interface{}) {
	log.std.Panicf(format, args...)
	log.err.Panicf(format, args...)
}

func (log *Logger) Fatal(args ...interface{}) {
	log.std.Fatal(args...)
	log.err.Fatal(args...)
}

func (log *Logger) Fatalf(format string, args ...interface{}) {
	log.std.Fatalf(format, args...)
	log.err.Fatalf(format, args...)
}

func (log *Logger) Writer() *io.PipeWriter {
	return log.std.Writer()
}
