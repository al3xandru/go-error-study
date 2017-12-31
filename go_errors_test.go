package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createError(msg string) error {
	return errors.New(msg)
}

func callError(wrap bool, msg string) (bool, error) {
	if err := createError(msg); err != nil {
		if wrap {
			// using fmt.Errorf is one of the approaches
			// to manually add sort of a trace
			return false, fmt.Errorf("wrap: %v", err)
		}
		return false, err
	}
	return true, nil
}

func TestErrorDetails(t *testing.T) {
	assert := assert.New(t)
	errorMsg := "main error message"
	_, err := callError(false, errorMsg)
	assert.Equal(errorMsg, err.Error(), "there is no stack trace")
}

func TestErrorFromFmt(t *testing.T) {
	assert := assert.New(t)
	errorMsg := "main error message"
	_, err := callError(true, errorMsg)
	assert.Equal("wrap: "+errorMsg, err.Error(), "there is not stack trace")

}
