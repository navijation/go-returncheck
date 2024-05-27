# go-returncheck
Go analyzer that requires function return values be used or explicitly ignored with `_`.

NOTE: This project is in active development. Documentation will not accurately reflect feature support until the first release.

# Usage

```shell
go install github.com/navijation/go-returncheck/cmd/returncheck@latest

returncheck [flags] [packages]
```

The `returncheck.Analyzer` variable follows guidelines in the golang.org/x/tools/go/analysis package. This should make it possible to integrate returncheck with your own analysis driver program.

# Motivation

In Go, functions that have no side-effects (so-called pure functions) can be called without storing or otherwise using their return values. This is almost always a bug. Other functions have side side-effects but also return values in which the caller should take careful interest. Accidentally disregarding these values _might_ be a bug. A common example of the former case is functions that return an updated value rather than mutating the original value.

```go
import (
  "context"

  "example.com/logging"
)

func UpdateContextLoggingField(ctx context.Context, field logging.Field) context.Context {
  // creates a copy of the logging fields to avoid mutating original context
  loggingFields := logging.ContextFields(ctx)
  loggingValues[field.Key] = field.Value
  return logging.ContextWithFields(ctx, loggingValues)
}

func main() {
  ctx := context.Background()

  // This function call is ineffectual
  UpdateContextLoggingFields(ctx, logging.NewField("caller", "main()"))

  // The logger won't emit "caller=main()" now
  logging.WithContext(ctx).Warn("Not yet implemented")
}
```

Callers can remove ambiguity about whether they are deliberately or accidentally ignoring return values by assigning return values to `_`. The [unusedresult analyzer](https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/unusedresult#hdr-Analyzer_unusedresult) enforces this unambiguous syntax for several well-known pure functions in the Go standard library. However, it is not extensible to cover any functions outside the standard library.

This analyzer, under the default configuration, reports **all** function calls (aside from some well-known exceptions in the standard library such as `fmt.Printf`) where the return values of the function call aren't explicitly assigned to `_`. This can improve the robustness of Go programs that contain functions which are not-so-obviously ineffectual unless their return values are used.

Future iterations of this analyzer will support allowlisting and blocklisting functions based on comment directives, as well as comment directives to indicate disallowing `_` assignments altogether for some functions.

# Example

```go
package main

func return5() int {
  return 5
}

func main() {
  return5()
  fmt.Println("Hello world!") 
}
```

Running returncheck with the default flags will emit the following:

```
main.go:8:2: function return5() return value is implicitly ignored
```

# Contributing

Contributions are welcome. Create PRs for existing issues or open up a new issue first.
