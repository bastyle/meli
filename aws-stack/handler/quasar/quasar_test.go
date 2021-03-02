package main

import "testing"

func TestMsg(t *testing.T) {
	if msg := getMessage(); (Message{}) != msg {
		t.Log("TestExpectedMsg PASSED")
	} else {
		t.Error("FAILED.")
	}
}
