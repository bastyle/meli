package messages

import (
	"testing"
)

var kenobi_msg = [5]string{"este", "", "", "mensaje", ""}
var skywalker_msg = [5]string{"", "es", "", "", "secreto"}
var sato_msg = [5]string{"este", "", "un", "", ""}
var kenobi_msg_6 = [6]string{"este", "", "", "mensaje", "", ""}
var expectedMsg = "este es un mensaje secreto"

func TestMsgLen(t *testing.T) {
	if valid, err := ValidateMessagesLen(kenobi_msg_6[:], skywalker_msg[:], sato_msg[:]); err == nil {
		t.Errorf("Expected: %v, got: %v", false, valid)
		t.Errorf("Se esperaba ser invalido debido al largo de uno de los mensajes.")
	} else {
		t.Log("TestMsgLen PASSED")
	}
	if _, err := ValidateMessagesLen(kenobi_msg[:], skywalker_msg[:], sato_msg[:]); err != nil {
		t.Errorf("TestMsgLen FAILED: All messages were expected to have the same number of words.")
	} else {
		t.Log("TestMsgLen 2 PASSED")
	}
}

func TestExpectedMsg(t *testing.T) {
	if msg := GetMessage(kenobi_msg[:], skywalker_msg[:], sato_msg[:]); msg != expectedMsg {
		t.Errorf("FAILED: Expected: %v, got: %v", expectedMsg, msg)
	} else {
		t.Log("TestExpectedMsg PASSED")
	}
}
