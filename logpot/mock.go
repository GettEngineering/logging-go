package logpot

func MockLog(logger Logger) {
	builder = func() Logger {
		return logger
	}
}

func UnMockLog() {
	builder = defaultBuilder
}
