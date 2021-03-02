package main

import "testing"

func TestMsg(t *testing.T) {
	if msg := getMessage(); msg != nil {
		t.Log("TestExpectedMsg PASSED")
	} else {
		t.Error("FAILED.")
	}
}
