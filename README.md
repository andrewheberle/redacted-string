# redacted-string

[![Go Report Card](https://goreportcard.com/badge/github.com/andrewheberle/redacted-string?style=flat-square)](https://goreportcard.com/report/github.com/andrewheberle/redacted-string)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/andrewheberle/redacted-string)](https://pkg.go.dev/mod/github.com/andrewheberle/redacted-string)

This is a very simple module that can be used to print a string as all '*' characters, although this is configurable.

## Install

```shell
go get github.com/andrewheberle/redacted-string
```

## Usage

Usage can be as simple as using the `Redact` call as follows:

```go
fmt.Println(redacted.Redact("a sensitive string"))
// Output: ********
```

Or you may customise the behaviour as follows:

```go
censor := redacted.New("a sensitive string", redacted.WithCharacter('!'), redacted.WithLength(12))
// Output: !!!!!!!!!!!!
```
