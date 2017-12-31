package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestingSkip(t *testing.T) {
	t.Skip(t.Name())
	assert.Fail(t, "Skip should never get to this step")
}

func TestTestingSkipf(t *testing.T) {
	t.Skipf("t.Skipf %s", t.Name())
	assert.Fail(t, "Skipf should never get to this step")
}

// Test will fail cause of `t.Error` and not due to the assertion
func TestTestingError(t *testing.T) {
	t.Error("t.Error")
	assert.True(t, t.Failed(), "expect %s marked as failed, actual %t", t.Name(), t.Failed())
}

// Test will fail cause of `t.Error` and not due to the assertion
func TestTestingErrorf(t *testing.T) {
	t.Errorf("firstly %s is not yet marked as failed %t", t.Name(), t.Failed())
	assert.True(t, t.Failed(), "expect %s marked as failed, actual %t", t.Name(), t.Failed())
}

func TestTestingFail(t *testing.T) {
	t.Fail()
	assert.True(t, t.Failed(), "expect %s marked as failed, actual %t", t.Name(), t.Failed())
}

func TestTestingFatal(t *testing.T) {
	t.Fatal("t.Fatal() on %s", t.Name())
	assert.Fail(t, "Fatal stops execution right away")
}

func TestTestingFatalf(t *testing.T) {
	t.Fatal("t.Fatal on %s", t.Name())
	assert.Fail(t, "t.Fatal stops execution right away")
}

func ExampleShowingOutputBasedTesting() {
	fmt.Println("this is it")
	// Output: this is it
}
