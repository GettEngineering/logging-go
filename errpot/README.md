# errArt - WIP

## Description:
TODO


## Notes:

1. When `err == nil`, `errpot.Wrap(err, "...")` will return `nil`, so `WithField/s(...)` can't be used

2. const errors should be created with `NewConstError(...)`
