## kisslog - keep it simple silly

Minimalistic logging library for Go, heavily inspired by [azer/logger](https://github.com/azer/logger) and borrowing from [chzyer/logex](https://github.com/chzyer/logex)
## Install

```bash
$ go get github.com/danwakefield/kisslog
```

## Manual

First create an instance with a preferred name.
The returned Logger has three methods, all of which act like `printf` style functions.
```go
package main

import "github.com/danwakefield/kisslog"

var log = kisslog.New("app")

func main() {
	log.Info("Requesting an image at %s", "foo/bar.jpg")
	log.Debug("I bet you wont see this")
	log.Error("Failed to start, shutting down")
}

//  [20:15:00][INFO ][main.main:simple.go:8] app: Requesting an image at foo/bar.jpg
//  [20:15:00][ERROR][main.main:simple.go:14] app: Failed to start, shutting down
```

Structured information can be added by passing a `kisslog.Attrs`
instance as the last variable to the logging function

```go
log.Info("I have just completed a task", kisslog.Attrs{
    "foo": 1,
    "bar": "baz",
})
// [20:18:33][INFO ][main.main:simple.go:12] app: I have just completed a task [ foo=1 bar=baz ]
```

### Options

The default settings for kisslog are to log Info and Error for every logger to stderr,
with output including the location of the log call.


#### Logging Level
By setting the `LOG_LEVEL` variable you can disable logging of certain methods.
The same can be achieved by setting `kisslog.LogLevel`

|       ENV VAR       |                  CODE                     |               INFO                 |
| ------------------- | ----------------------------------------- | ---------------------------------- |
| `LOG_LEVEL=DEBUG`   | `kisslog.LogLevel = kisslog.DebugLevel`   | Allows `Debug`, `Info` and `Error` |
| `LOG_LEVEL=INFO`    | `kisslog.LogLevel = kisslog.InfoLevel`    | Allows `Info` and `Error`          |
| `LOG_LEVEL=ERROR`   | `kisslog.LogLevel = kisslog.ErrorLevel`   | Allows `Error`                     |
| `LOG_LEVEL=DISABLE` | `kisslog.LogLevel = kisslog.DisableLevel` | Allows no logging                  |

#### Filtering
Loggers can be filtered by setting the `LOG_ENABLED` variable.
It can be managed programmatically using `kisslog.EnableLogger` and `kisslog.DisableLogger`.
Using any of these options disables logging from a logger unless it is explicitly enabled.

|       ENV VAR       |           CODE                 |               INFO                                  |
| :-----------------: | ------------------------------ | --------------------------------------------------- |
| `LOG_ENABLED=foo`   | `kisslog.EnableLogger("foo")`  |  All loggers not 'enabled' explicitly are disabled  |
|        NA           | `kisslog.DisableLogger("foo")` |  Loggers cannot be disabled with ENVARS             |


#### JSON Output
JSON Output can be forced on and off by setting the `LOG_JSON` variable.
The `kisslog.JSONOutput` variable can be used for programmatic control.

If the variable is not set JSON output is used by default if stderr is
not a terminal.

|       ENV VAR       |           CODE               |               INFO                                  |
| ------------------- | ---------------------------- | --------------------------------------------------- |
| `LOG_JSON=TRUE`     | `kisslog.JSONOutput = true`  |  Can also use `1`, `on` or `enable` to set ENVVAR   |
| `LOG_JSON=FALSE`    | `kisslog.JSONOutput = false` |  Can also use `0`, `off` or `disable` to set ENVVAR |

```
LOG_JSON=TRUE go run example/simple.go
# {"time":"20:59:35","package":"app","level":"[INFO ]","trace":"[main.main:simple.go:8] ","msg":"Requesting an image at foo/bar.jpg"}
# {"time":"20:59:35","package":"app","level":"[INFO ]","trace":"[main.main:simple.go:12] ","msg":"I have just completed a task","attributes":{"bar":"baz","foo":1}}
# {"time":"20:59:35","package":"app","level":"[ERROR]","trace":"[main.main:simple.go:14] ","msg":"Failed to start, shutting down"}
```

#### Function Tracing
By default kisslog adds the location of the log call to its output using runtime inspection.
This can be disabled using `LOG_TRACE` or `kisslog.TraceFile`

|       ENV VAR       |           CODE               |               INFO                                  |
| ------------------- | ---------------------------- | --------------------------------------------------- |
| `LOG_TRACE=TRUE`    | `kisslog.TraceFile = true`   |  Can also use `1`, `on` or `enable` to set ENVVAR   |
| `LOG_TRACE=FALSE`   | `kisslog.TraceFile = false`  |  Can also use `0`, `off` or `disable` to set ENVVAR |

```
LOG_TRACE=FALSE go run example/simple.go
# [16:23:37][INFO ]app: Requesting an image at foo/bar.jpg
# [16:23:37][INFO ]app: I have just completed a task [ foo=1 bar=baz ]
# [16:23:37][ERROR]app: Failed to start, shutting down
```

#### Time format
You can change the precision of the timestamp using `LOG_TIMEFORMAT` or `kisslog.TimeFormat`.
These both take values that [golang's time package](https://golang.org/pkg/time/#Constants) can parse.

|       ENV VAR                              |           CODE                      |
| ------------------------------------------ | ----------------------------------- |
| `LOG_TIMEFORMAT=2006-01-02T15:04:05Z07:00` | `kisslog.TimeFormat = time.RFC3339` |

```
LOG_TIMEFORMAT=2006-01-02T15:04:05Z07:00 go run example/simple.go
# [2000-12-20T20:19:26+01:00][INFO ][main.main:simple.go:8] app: Requesting an image at foo/bar.jpg
# [2000-12-20T20:19:26+01:00][INFO ][main.main:simple.go:12] app: I have just completed a task [ foo=1 bar=baz ]
# [2000-12-20T20:19:26+01:00][ERROR][main.main:simple.go:14] app: Failed to start, shutting down
```
#### Output Stream
By default kisslog logs to `os.Stderr`.
This can be changed by calling `kisslog.SetOutput()` with an `io.Writer`

##### Output logs to a new file
```go
f, err := os.Create("foo.log")
if err != nil {
    panic(err)
}
defer f.Close() // Ensure that the file is closed.
kisslog.SetOutput(f)
```

##### Append logs to existing file
```go
f, err := os.OpenFile("foo.log", os.O_APPEND|os.O_WRONLY, 0666)
if err != nil {
    panic(err)
}
defer f.Close()
kisslog.SetOutput(f)
```

##### Append if file exists otherwise create it
```go
f, err := os.OpenFile("foo.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
if err != nil {
    panic(err)
}
defer f.Close()
kisslog.SetOutput(f)
```
