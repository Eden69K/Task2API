package logger

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	SetOutput(w io.Writer)
}

type StdLogger struct {
	logger *log.Logger
}

//------------------------------------------------------------------------------------------

func NewStdLogger() *StdLogger {
	return &StdLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

//------------------------------------------------------------------------------------------

func (l *StdLogger) Debug(format string, args ...interface{}) {
	l.logger.Printf("[DEBUG] "+format, args...)
}

func (l *StdLogger) Info(format string, args ...interface{}) {
	l.logger.Printf("[INFO] "+format, args...)
}

func (l *StdLogger) Warn(format string, args ...interface{}) {
	l.logger.Printf("[WARN] "+format, args...)
}

func (l *StdLogger) Error(format string, args ...interface{}) {
	l.logger.Printf("[ERROR] "+format, args...)
}

//------------------------------------------------------------------------------------------

func (l *StdLogger) SetOutput(w io.Writer) {
	l.logger.SetOutput(w)
}

func ConfigureFileLogger(logger *StdLogger, filename string, maxSizeMB, maxBackups, maxAgeDays int) {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSizeMB,
		MaxBackups: maxBackups,
		MaxAge:     maxAgeDays,
		Compress:   true,
	}

	multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)
	logger.SetOutput(multiWriter)
}
