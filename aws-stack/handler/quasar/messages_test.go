package main

import (
	"testing"
)

func TestValidateMessagesLen(t *testing.T) {
	if _, err := ValidateMessagesLen(kenobi_msg[:], skywalker_msg[:], sato_msg[:]); err != nil {
		t.Errorf("TestMsgLen FAILED: All messages were expected to have the same number of words.")
	}
	if valid, err := ValidateMessagesLen(kenobi_msg_6[:], skywalker_msg[:], sato_msg[:]); err == nil {
		t.Errorf("Expected: %v, got: %v", false, valid)
	}
}

func TestGetMessage(t *testing.T) {
	if msg := GetMessage(kenobi_msg[:], skywalker_msg[:], sato_msg[:]); msg != expectedMsg {
		t.Errorf("FAILED: Expected: %v, got: %v", expectedMsg, msg)
	}
}
