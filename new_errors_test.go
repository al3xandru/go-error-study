package main

import (
	"bytes"
	"fmt"
	newerr "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createOrWrapError(msg string, cause error) (bool, error) {
	if cause == nil {
		return true, fmt.Errorf(msg)
	} else {
		return false, newerr.Wrap(cause, msg)
	}
}

func TestWrap(t *testing.T) {
	assert := assert.New(t)
	errorMsg := "root error message"
	_, err := createOrWrapError(errorMsg, nil)
	assert.Equal(errorMsg, err.Error(), "root error should preserve the message")

	_, err = createOrWrapError("wrap", err)
	assert.Equal(fmt.Sprintf("%s: %s", "wrap", errorMsg), err.Error(), "wrapped error concatenates messages")
}

func TestCause(t *testing.T) {
	assert := assert.New(t)
	_, err := createOrWrapError("root", nil)
	_, werr := createOrWrapError("lvl2", err)

	assert.NotNil(newerr.Cause(err), "cause of root error is not empty")
	assert.Equal(err, newerr.Cause(err), "cause of root error is the error itself")
	assert.NotNil(newerr.Cause(werr), "wrapped errors have a cause")
	assert.Equal(err, newerr.Cause(werr))
}

func TestErrorsErrorf(t *testing.T) {
	assert := assert.New(t)
	err := newerr.Errorf("my bad: %s", "root error")
	assert.Equal("my bad: root error", err.Error())
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%+v", err)
	out := buf.String()
	assert.Contains(out, err.Error(), "expected logged error to contain error message '%s'", err.Error())
	assert.Contains(out, "new_errors_test.go:", "expected logged error to contain stack trace")
}

type stackTracer interface {
	StackTrace() newerr.StackTrace
}

func TestErrorsNew(t *testing.T) {
	assert := assert.New(t)
	err := newerr.New("root error")
	wrappedErr, ok := err.(stackTracer)
	assert.True(ok, "errors.New expected to return an error supporting StackTrace()")
	assert.True(len(wrappedErr.StackTrace()) > 0, "errors.New stack trace expected > 0 found %d", len(wrappedErr.StackTrace()))
}

func TestWithStack(t *testing.T) {
	_, err := createOrWrapError("root", nil)
	err = newerr.WithStack(err)
	wrappedErr, ok := err.(stackTracer)
	assert.True(t, ok, "errors.New expected to return an error supporting StackTrace()")
	assert.True(t, len(wrappedErr.StackTrace()) > 0, "errors.New stack trace expected > 0 found %d", len(wrappedErr.StackTrace()))
}

func ExampleWithMessage() {
	rootErr := newerr.New("root")
	wrapErr := newerr.WithMessage(rootErr, "wrap")
	fmt.Println(wrapErr)
	// Output: wrap root
}
