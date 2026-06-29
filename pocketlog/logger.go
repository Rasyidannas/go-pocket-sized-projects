package pocketlog

// Logger is used to log information
type Logger struct {
	threshold Level
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any){
	// implement me
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	// implement me
}

// New returns you a logger, ready to log at the required threshold.
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}
