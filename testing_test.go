package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSkip(t *testing.T) {
	t.Skip(t.Name())
	assert.Fail(t, "Skip should never get to this step")
}
func TestSkipf(t *testing.T) {
	t.Skipf("t.Skipf %s", t.Name())
	assert.Fail(t, "Skipf should never get to this step")
}

func TestError(t *testing.T) {
	t.Error("t.Error")
	t.Logf("test %s passed %t", t.Name(), !t.Failed())
}

func TestErrorf(t *testing.T) {
	t.Errorf("firstly %s is not yet marked as failed %t", t.Name(), t.Failed())
	t.Logf("test %s passed %t", t.Name(), !t.Failed())
}

func TestFail(t *testing.T) {
	t.Fail()
	assert.True(t, t.Failed(), "%s was expected to be marked as failed", t.Name())
}

func TestFatal(t *testing.T) {
	t.Fatal("t.Fatal")
	assert.Fail(t, "Fatal stops execution right away")
}

func ExampleShowingOutputBasedTesting() {
	fmt.Println("this is it")
	// Output: this is not it
}
