package main

import (
	"testing"
)

var kenobi_msg = [5]string{"este", "", "", "mensaje", ""}
var skywalker_msg = [5]string{"", "es", "", "", "secreto"}
var sato_msg = [5]string{"este", "", "un", "", ""}
var kenobi_msg_6 = [6]string{"este", "", "", "mensaje", "", ""}
var expectedMsg = "este es un mensaje secreto"

func TestMsg(t *testing.T) {
	if msg := getMessage(); (Message{}) != msg {
		t.Log("TestExpectedMsg PASSED")
	} else {
		t.Error("FAILED.")
	}
}

func TestExpectedMsg(t *testing.T) {
	//if msg := GetMessage(kenobi_msg[:], skywalker_msg[:], sato_msg[:]); msg != expectedMsg {
	if msg := GetExampleMessage(); msg.Message != expectedMsg {
		t.Errorf("FAILED: Expected: %v, got: %v", expectedMsg, msg)
	} else {
		t.Log("TestExpectedMsg PASSED")
	}
}
