package errpot

type ErrPot interface {
	WithField(key string, val interface{}) ErrPot
	WithFields(fields Fields) ErrPot
	WithFieldsFrom(holder FieldsHolder) ErrPot
	Error() string
}

type ErrorWithFields interface {
	GetError() error
	GetFields() map[string]interface{}
}

type FieldsHolder interface {
	GetFields() map[string]interface{}
}

type Fields map[string]interface{}
