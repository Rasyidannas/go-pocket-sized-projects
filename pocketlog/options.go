package pocketlog

import "io"

// Option define a functional option to our logger
type Option func(*Logger)

// WithOutput returns a configuration function that sets the output of logs.
func WithOutput(output io.Writer) Option {
	return func(lgr *Logger) {
		lgr.output = output
	}
}
