package logger

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"time"
)

// Logger defines the structure of our custom logger
type Logger struct {
	info  *color.Color
	warn  *color.Color
	error *color.Color
	debug *color.Color
	level int
}

// New creates a new instance of Logger
func New() *Logger {
	return &Logger{
		info:  color.New(color.FgGreen),
		warn:  color.New(color.FgYellow),
		error: color.New(color.FgRed),
		debug: color.New(color.FgBlue),
		level: 2,
	}
}

// Info logs an info message
func (l *Logger) Info(format string, a ...interface{}) {
	l.log(l.info, "INFO", format, a...)
}

// Warn logs a warning message
func (l *Logger) Warn(format string, a ...interface{}) {
	l.log(l.warn, "WARN", format, a...)
}

// Error logs an error message
func (l *Logger) Error(format string, a ...interface{}) {
	l.log(l.error, "ERROR", format, a...)
}

// Debug logs a debug message
func (l *Logger) Debug(format string, a ...interface{}) {
	l.log(l.debug, "DEBUG", format, a...)
}

// log is the internal logging function
func (l *Logger) log(c *color.Color, level string, format string, a ...interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	message := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "%s [%s] %s\n", timestamp, level, c.Sprint(message))
}
