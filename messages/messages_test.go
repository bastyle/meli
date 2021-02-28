package messages

import (
	"testing"
)

var kenobi_msg = [5]string{"este", "", "", "mensaje", ""}
var skywalker_msg = [5]string{"", "es", "", "", "secreto"}
var sato_msg = [5]string{"este", "", "un", "", ""}
var kenobi_msg_6 = [6]string{"este", "", "", "mensaje", "", ""}

func TestMsgLen(t *testing.T) {
	t.Log("TestMsgLen...")
	if valid, err := ValidateMessagesLen(kenobi_msg_6[:], skywalker_msg[:], sato_msg[:]); err == nil {
		t.Errorf("Expected: %v, got: %v", false, valid)
		t.Errorf("Se esperaba ser invalido debido al largo de uno de los mensajes.")
	}
	if valid, err := ValidateMessagesLen(kenobi_msg[:], skywalker_msg[:], sato_msg[:]); err != nil {
		t.Errorf("Expected: %v, got: %v", true, valid)
		t.Errorf("Se esperaba que todos los mensajes tuviera la misma cantidad de palabras.")
	}
}
