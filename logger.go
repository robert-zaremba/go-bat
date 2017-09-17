package bat

// Logger interface
type Logger interface {
	Error(msg string, ctx ...interface{})

	// Fatal is a Crit log followed by panic
	Fatal(msg string, ctx ...interface{})
}
