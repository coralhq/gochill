# Gochill
Golang library for building logging data structure.

# Abstract

Logging data structure that supported is following [GELF](http://docs.graylog.org/en/2.0/pages/gelf.html#).
All data structure should be in json mode.  Required fields :

- version
- host
- service
- short_message
- timestamp
- level

We need a library that support this logging data structure but also should be easy
to use for other programmers.

# Rationale

This library should be only for building logging data structure and print it using
stdout or stderr.

# Environment Variable

This library will listen to environment variable with registered key `SERVICE_NAME`.
This variable used to fill `service` json field.

If environment variable not exists, then `service` will be an empty string.

Log output format have 2 type `JSON` and `TEXT`, if use `TEXT` it will use library [Logrus](https://github.com/sirupsen/logrus) for 
output. this can be set to environment variable with registered key `JSON_FORMAT`, this boolean value.
if set false it will use `TEXT` if true output will set to `JSON` format. default value always set to true

# Hostname

This library will automatically load `hostname` from os using golang's native library

```go
os.Hostname()
```

But if there is some error, then `hostname` will be an empty string.

# Usage

```go
import "gochill"

gochill.Alert(Msg("alert message"))
```

Available levels :

- Alert
- Critical
- Error
- Warning
- Notice
- Info
- Debug

Gochill designed with `io.MultiWriter` it means you can add your own custom
`io.Writer` to proceed your log.

```go
package main

import "gochill"

type CustomWriter struct{}
func (cw CustomWriter) Write(message []byte) (int, error) {
  fmt.Println("hello world")
}

gochill.CustomOutput = CustomWriter{}
gochill.Info(Msg("log info"))
```

These code should produce an output will send a structured json code and also
print "hello world" to stdout.

Each of log function also support for optional fields, usages :

```go
package main

import "gochill"

gochill.Alert(Msg("log alert"), O("key1", "value1"), O("key2", 230))
```

It will produce result like this :

```
{"_key1":"value1","_key2":230,"host":"your.hostname","level":1,"service":"","short_message":"log alert","timestamp":1472879389234,"version":1}
```

If you need a `full_message` to describe your log, you can filled optional full message like this :

```go
package main

import "gochill"

gochill.Alert(Msg("short message", "full_message"))
```

# Unit Test

How to run :

```
go test -cover
```

Via `Makefile` :

```
make test
```
