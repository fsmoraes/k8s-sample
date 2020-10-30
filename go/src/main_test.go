package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	message := "hello"
	want := "<b>hello</b>"
	result := greeting(message)

	if want != result {
		t.Fail()
	}
}

func TestEmptyMessage(t *testing.T) {
	message := ""
	want := "<b></b>"
	result := greeting(message)

	if want != result {
		t.Fail()
	}
}
