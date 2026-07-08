package pocketlog

import (
	"fmt"
	"os"
)

// Logger is used to log information
type Logger struct {
	threshold Level
	output io.Writer
}

// Debugf formats and prints a message if the log level is debug or lower.
func (l *Logger) Debugf(format string, args ...any){
	// making sure we can safely write to the output
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelDebug{
		_, _ = fmt.Fprintf(l.output, format+"\n", args...)
	}
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
// The default output is Stdout
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
		output: output,
	}
}
