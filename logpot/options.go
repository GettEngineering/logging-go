package logpot

type LogOptions struct {
	PrintFieldsInsideMessage bool
	PrintAsFields            []string
	PrintErrorWithStackTrace bool
}

var DefaultLogOptions = LogOptions{
	PrintFieldsInsideMessage: false,
	PrintAsFields:            []string{"error"},
	PrintErrorWithStackTrace: false,
}

var logOptions = DefaultLogOptions

func SetLogOptions(options LogOptions) {
	logOptions = options
}
