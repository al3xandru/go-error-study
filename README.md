# Error handling in Go

Read the first 3 links in the Resources section.  They summarize both the space
of potential problems handling errors and offer the correct alternatives.

There are some Go libraries that are adding more context to errors on top of the
default Go error system.  Being able to wrap errors and have stack traces can
significantly improve the debuggability.

This repo is experimenting with the <https://godoc.org/github.com/pkg/errors> API.

## Projects

* <https://godoc.org/github.com/pkg/errors>
* <http://github.com/hashicorp/go-multierror>
* <http://github.com/goware/errorx>

## Resources

* [Error handling and Go - The Go Blog](https://blog.golang.org/error-handling-and-go)
* [Don’t just check errors, handle them gracefully | Dave Cheney](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
* [Error handling patterns in Go](https://mijailovic.net/2017/05/09/error-handling-patterns-in-go/)
* [Advanced Error Handling in Golang](http://blog.ralch.com/articles/advanced-error-handling-in-golang/)
* [Go Best Practices — Error handling – Sebastian Dahlgren – Medium](https://medium.com/@sebdah/go-best-practices-error-handling-2d15e1f0c5ee)

* * * * *

# Running tests

```
go test -run ''         # Run all tests
go test -run Foo        # Run top-level tests matching "Foo" (e.g.  "TestFooBar")
```

The argument to `-run` and `-bench` command-line flags is a regular expression
that matches the test's name.  This also enables [running subtests](https://golang.org/pkg/testing/#hdr-Subtests_and_Sub_benchmarks):

```
go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
go test -run /A=1    # For all top-level tests, run subtests matching "A=1".
```

The `-v` flag prints the name and execution time of each test in the package.

# TODO

- [ ] Improve examples in `go_errors_test.go` to capture the lack of context of
    Go errors
- [ ] Add tests in `go_errors_test.go` to exemplify the usage of sentinel
    errors, error types, and opaque errors
- [ ] Can tests in `testing_test.go` to make them pass

# License

Code in this repo is licensed under Apache License v2.0.

Copyright Alex Popescu 2016-2017
