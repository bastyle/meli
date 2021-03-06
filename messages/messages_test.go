package messages

import (
	"testing"
)

const kMsg = [5]string{"este", "", "", "mensaje", ""}
const skMsg = [5]string{"", "es", "", "", "secreto"}
const saMsg = [5]string{"este", "", "un", "", ""}
const kMsg_6 = [6]string{"este", "", "", "mensaje", "", ""}

const exMsg = "este es un mensaje secreto"

func TestMsgLen(t *testing.T) {
	if valid, err := ValidateMessagesLen(kMsg_6[:], skMsg[:], saMsg[:]); err == nil {
		t.Errorf("Expected: %v, got: %v", false, valid)
		t.Errorf("Se esperaba ser invalido debido al largo de uno de los mensajes.")
	} else {
		t.Log("TestMsgLen PASSED")
	}
	if _, err := ValidateMessagesLen(kMsg[:], skMsg[:], saMsg[:]); err != nil {
		t.Errorf("TestMsgLen FAILED: All messages were expected to have the same number of words.")
	} else {
		t.Log("TestMsgLen 2 PASSED")
	}
}

func TestExpectedMsg(t *testing.T) {
	if msg := GetMessage(kMsg[:], skMsg[:], saMsg[:]); msg != exMsg {
		t.Errorf("FAILED: Expected: %v, got: %v", exMsg, msg)
	} else {
		t.Log("TestExpectedMsg PASSED")
	}
}
