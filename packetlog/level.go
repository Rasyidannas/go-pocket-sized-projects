package packetlog

type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used for debugging
	LevelDebug Level = iota // this iota will allows us to create a sequence of numbers increments
	// LevelInfo represents a logging level that contains information deemed
	LevelInfo
	// LevelError represents the highest logging level, only to be used to tr
	LevelError
)
