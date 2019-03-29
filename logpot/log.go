package logpot

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type log struct {
	fields Fields
	err    error
}

var defaultBuilder = func() Logger {
	return log{
		fields: Fields{},
		err:    nil,
	}
}

var builder = defaultBuilder

// log constructors ///////////////////////////////////////////////////////////

func WithField(key string, value interface{}) Logger {
	x := builder().WithField(key, value)
	return x
}

func WithFields(fields Fields) Logger {
	return builder().WithFields(fields)
}

func WithError(err error) Logger {
	return builder().WithError(err)
}

func Debug(args ...interface{}) {
	builder().Debug(fmt.Sprint(args...))
}

func Info(args ...interface{}) {
	builder().Info(fmt.Sprint(args...))
}

func Warn(args ...interface{}) {
	builder().Warn(fmt.Sprint(args...))
}

func Error(args ...interface{}) {
	builder().Error(fmt.Sprint(args...))
}

// log builders /////////////////////////////////////////////////////////////////

func (l log) WithField(key string, value interface{}) Logger {
	l.fields[key] = fmt.Sprintf("%+v", value) // %+v print complex objects
	return l
}

func (l log) WithFields(fields Fields) Logger {
	for k, v := range fields {
		l.fields[k] = v
	}
	return l
}

func (l log) WithError(err error) Logger {
	// if error is fields provider
	if e, ok := err.(ErrorWithFields); ok {
		// get error
		err = e.GetError()

		// copy error fields to current log
		for k, v := range e.GetFields() {
			l.fields[k] = v
		}
	}
	l.err = err
	return l
}

func (l log) GetFields() map[string]interface{} {
	return l.fields
}

// prints /////////////////////////////////////////////////////////////////

func (l log) Debug(msg string) {
	l.addErrorIfExist()
	asFields, asMsg := segregateFields(l.fields)
	logrus.WithFields(logrus.Fields(asFields)).Debug(msg, asMsg.toString())
}

func (l log) Info(msg string) {
	l.addErrorIfExist()
	asFields, asMsg := segregateFields(l.fields)
	logrus.WithFields(logrus.Fields(asFields)).Info(msg, asMsg.toString())
}

func (l log) Warn(msg string) {
	l.addErrorIfExist()
	asFields, asMsg := segregateFields(l.fields)
	logrus.WithFields(logrus.Fields(asFields)).Warn(msg, asMsg.toString())
}

func (l log) Error(msg string) {
	l.addErrorIfExist()
	asFields, asMsg := segregateFields(l.fields)
	logrus.WithFields(logrus.Fields(asFields)).Error(msg, asMsg.toString())
}

func (l *log) addErrorIfExist() {
	if l.err != nil {
		if logOptions.PrintErrorWithStackTrace {
			l.fields["error"] = l.err
		} else {
			l.fields["error"] = l.err.Error()
		}
	}
	return
}
