package main_test

import (
	a "codenut.org/lab/go/examples/test"
	"testing"
)

var subject = &a.A{}

func TestSuccess(t *testing.T) {
	expected := "something"
	x := subject.DoSomething()
	if x != expected {
		t.Fatal("Expected '" + expected + "' but got '" + x + "'")
	}
}

func TestFailure(t *testing.T) {
	expected := "nope"
	x := subject.DoSomething()
	if x != expected {
		t.Fatal("Expected '" + expected + "' but got '" + x + "'")
	}
}
