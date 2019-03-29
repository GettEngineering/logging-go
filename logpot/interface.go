package logpot

//go:generate mockgen -source=./logpot/interface.go  -destination=./logpot/mock/logpot_mock.go Logger
type Logger interface {
	WithField(key string, value interface{}) Logger
	WithFields(fields Fields) Logger
	WithError(err error) Logger

	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)

	// This method allows transparently return log's fields.
	// Could be useful when log is passed to entity that
	// might process log's data (e.g. error processor that could
	// use passed log's data to add to the error message.
	GetFields() map[string]interface{}
}

// This interface allows deal with smart errors
// Could be useful when log should print error that implements
// this interface (means have in addition to error also map of fields)
// See usage in code for more details.
type ErrorWithFields interface {
	GetError() error
	GetFields() map[string]interface{}
}
