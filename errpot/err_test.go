package errpot_test

import (
	"testing"
	"time"

	"github.com/gtforge/logging-go/errpot"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ErrorWithFields suit")
}

var _ = Describe("ErrorWithFields", func() {

	It("Error comparison", func() {
		constErr := errpot.NewConstError("err")

		err := func() error {
			return constErr
		}()

		same := err == constErr
		Expect(same).To(BeTrue())
		Expect(err).To(Equal(constErr))
	})

	It(".Error", func() {
		err1 := errpot.New("err_msg_1")
		err2 := errpot.Wrap(err1, "err_msg_2").WithField("f", "v")
		Expect(err2.Error()).To(Equal("err_msg_2: err_msg_1 {\"f\":\"v\"}"))
	})

	It("Wrap nil error", func() {
		err2 := errpot.Wrap(nil, "err_msg_2")
		Expect(err2).To(BeNil())
	})

	It("Basic usage", func() {

		// call function that might (and will) return error
		err := func() error {

			fieldsHolder := fieldsHolder{
				fields: map[string]interface{}{
					"field1": "val1",
					"field2": 2,
					"field3": true,
				},
			}

			// call another function that might (and will) returns error
			err := func() error {
				return errpot.New("most_internal_err_msg")
			}()

			// regular error check
			if err != nil {
				return errpot.Wrap(err, "something went wrong").
					WithFieldsFrom(fieldsHolder).
					WithField("field4", "val4").
					WithFields(errpot.Fields{
						"field5": "val5",
					})
			}

			return nil // will never happened
		}()

		logpot.WithError(err).Error("log_error")

		time.Sleep(100 * time.Millisecond)
	})

})

type fieldsHolder struct {
	fields map[string]interface{}
}

func (fp fieldsHolder) GetFields() map[string]interface{} {
	return fp.fields
}
