package pocketlog

import (
	"fmt"
)

// Logger is used to log information
type Logger struct {
	threshold Level
}

// Debugf formats and prints a message if the log level is debug or lower.
func (l *Logger) Debugf(format string, args ...any){
	if l.threshold > LevelDebug {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

// Infof formats and prints a message if the log level is info or lower.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

// Errorf formats and prints a message if the log level is error or lower
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

// Error logs an empty message at the error level.
func (l *Logger) Error() {
	l.Errorf("")
}

// New returns you a logger, ready to log at the required threshold.
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}
