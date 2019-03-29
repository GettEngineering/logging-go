# logpot


This package allows a programmers with different skills and tastes make
logging easier by providing limited and simple set of logging options.
On the other hand it gives an ability to define (by service owner or architect)
the standard log output.

-----

#### Why `WithField/s` - log management systems

Existing set of logging tools (e.g. logrus) allows programmers do the
following:
```
log.Debugf("action %s was done by %s", action, user)
log.Debugf("%s / %s", action, user)
log.Debugf("user: %s, action: %s", user, action)
```
All 3 outputs will contain exactly the same information with totally different shape.
When couple of people are working on the same code base the log's shapes became
more various that make it harder to consume.
This can be solved by defined (between programmers) conventions, but it
cannot be enforced by compiler and require additional resources (to define the
convention, to ensure it's not broken by new code, to explain to new comers
why this is the right way).
In addition, there is a lot of log management systems (ELK, Graylog, ...) that require
JSON - styled log data.

To provide JSON shaped logs previous examples became to:
```
logrus.WithField("action", action).Debugf("action is done by user %v", user)
logrus.WithField("user", user).Debugf("action %v is done")
logrus.WithField("action", action).WithField("user", user).Debug("action is done")
```
First two options have part of the data inside the log message and part
as fields. What is the difference? Some log management systems point
of view (e.g. ELK), will store the fields data (key and value) indexed to provide
faster search by the key.

In addition, first two options have same problem explained before - not standard
log shape. This also could be solved, but let's looks on the third option.

Third option includes variable data as fields and constant action description as
a log message. This looks good:

- Divide the action description and action data
- Make all logs look pretty same
- Well integrated with log management systems
- No conventions should be done

This is the reason current package doesn't support `Debug/Info/... + f`
If you need ad some data to the log - use `WithField/s`.

-----

#### Why `WithField/s` - log formatter for human readable logs

Sometimes we want read the logs without using log management systems.
For example during development, local testing or on test/stage environments.

Is it became clear after pretty short time that reading white colored log on black
screen (or black on white) is not the most convenient way.
It's much clearer to read colored and well shaped log.
Using `WithField/s` allows to log formatter do it's job much better.
What does it mean? To more information please refer to `logrus-human-formatter`
package.

-----

#### Why `WithField/s` - more

Logging can be seen as providing kind of context or state of current action.
From the examples above we can see what happened: "action is done" and the
context: "user", "action name".

This context can be used not only for logging, but also for the "context
based" activity.
For example, error happened during the action. So `err` will be returned
from current scope. It's make sense to add to this error pretty same data
as we already have in the log ("context").

To make this possible, all the data must be in fields. Otherwise there is
no "context".

\* this is little abstract reason to use `WithField/s`, but there is a real
usage - wip.

-----

#### Why not `WithField/s`

If all the the data will be added "as field", log management systems will
perform a lot of work to build the indexes to support search by the field's key.
This is not always necessary, for example we want have ability to fast search by
"action", but we never searching by "user".

In this case adding `...WithField("user", user)` is not desired.

-----

#### Summary

So we see why using `WithField/s` can be good and why can be not.
The $1M question is what should the programmer do. Use it? Not use it?
Use it sometimes? Use it for important for search keys? Looks like there is
no right answer `logrus` is providing.

The `logpot` package has simple solution for this dilemma:

- Always use `WithField/s` in the code. Focus on the code and not on the
log. Write log always in the same way. Simple, don't think - do.

- Define the important keys in one place (during code initialization or make it
dynamical). Each log record will pass though the segregation level that will
place the data as a field (if defined) or as the part of log message.
See `PrintFieldsInsideMessage`, `PrintAsFields` settings usage for more info.

## Usage

* Very similar to logrus (without option to `Debug/Info/...` + `f`)

```
logpot.WithField("field1", "value1").Debug("my log message")

logpot.WithFields(logpot.Fields{
    "field2": "value2",
    "field3": "value3",
}).Info("my log message")

logpot.WithField("field1", "value1").WithError(err).Error("my log message")
```

## LogOptions

- `PrintFieldsInsideMessage bool` - if true, all the fields (except of defined in `PrintAsFields`) will
    be printed as part of the log message.

- `PrintAsFields []string` - see previous

    For example, for these options:
    ```
    PrintFieldsInsideMessage = true
    PrintAsFields = []string{"field2"}
    ```

    This log:
    ```
    logpot.
      WithField("field1", "value1").
      WithField("field2", "value2").
      Debug("my log message")
    ```
    Will provide the same log as:
    ```
    logrus.
      WithField("field1", "value1").
      Debug("my log message {field2: value2}")
    ```

- `PrintErrorWithStackTrace bool` - define whether stuck trace is needed or no in the error log

See `log_test` file for some usages.


## Mock
Sometimes we want test functionality that doesn't return data.
In case it prints a log, with logpot it's possible to replace
original log with mock.

For example:
```
// before each:
mockCtrl = gomock.NewController(GinkgoT())
mockLog = mock_logpot.NewMockLogger(mockCtrl)
logpot.MockLog(mockLog) // <---- replace log with mock

// after each:
mockCtrl.Finish()

// expectaions:
mockLog.EXPECT().WithField("k1", "v1").Return(mockLog) // <---- returns mock
mockLog.EXPECT().WithField("k2", "v2").Return(mockLog)
mockLog.EXPECT().Debug("deb")

// somewhere in you code:
...
logpot.
    WithField("k1", "v1").
    WithField("k2", "v2").
    Debug("deb")
...
```
You can see this example in test.

