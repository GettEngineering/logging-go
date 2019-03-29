package errpot

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func Wrap(err error, msg string) ErrPot {

	if err == nil {
		return nil // pay attention, can't use WithField/s if err == nil
	}

	errPot := errPot{
		fields: Fields{},
	}

	// copy fields and error from input error
	if e, ok := err.(ErrorWithFields); ok {
		err = e.GetError()
		errPot = errPot.addFields(e.GetFields())
	}

	errPot.err = errors.Wrap(err, msg)
	return errPot
}

func NewConstError(msg string) error {
	return errors.New(msg)
}

func New(msg string) ErrPot {
	return errPot{
		fields: Fields{},
		err:    errors.New(msg),
	}
}

// container to pass error and fields as error object
type errPot struct {
	fields Fields
	err    error
}

func (r errPot) WithFieldsFrom(holder FieldsHolder) ErrPot {
	for k, v := range holder.GetFields() {
		r.fields[k] = v
	}
	return r
}

func (r errPot) WithFields(fields Fields) ErrPot {
	for k, v := range fields {
		r.fields[k] = v
	}
	return r
}

func (r errPot) addFields(fields Fields) errPot {
	for k, v := range fields {
		r.fields[k] = v
	}
	return r
}

func (r errPot) WithField(key string, val interface{}) ErrPot {
	r.fields[key] = fmt.Sprintf("%+v", val) // pretty print complex objects
	return r
}

func (r errPot) GetError() error {
	return r.err
}

func (r errPot) GetFields() map[string]interface{} {
	return Fields(r.fields)
}

// Implements error interface, so can be returned as error
func (r errPot) Error() string {
	return fmt.Sprint(r.err, " ", map2String(r.fields))
}

func map2String(m map[string]interface{}) string {
	// json.Marshal impl without using json.Marshal
	var pairs []string
	for k, v := range m {
		pairs = append(pairs, fmt.Sprintf("\"%v\":\"%+v\"", k, v))
	}
	return fmt.Sprintf("{%v}", strings.Join(pairs, ","))
}
