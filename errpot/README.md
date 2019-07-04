# errpot

## Description:

Very convenient if used together with a [logpot](https://github.com/GettEngineering/logging-go/edit/master/logpot):

```
log := logpot.WithFields(...)
...
errpot.Wrap(err, "error msg").WithFieldsFrom(log) // <-- no need add context data, use log's context 
```


## Notes:

1. When `err == nil`, `errpot.Wrap(err, "...")` will return `nil`, so `WithField/s(...)` can't be used

2. const errors should be created with `NewConstError(...)`
