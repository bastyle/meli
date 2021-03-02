package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*var kenobi_msg = [5]string{"este", "", "", "mensaje", ""}
var skywalker_msg = [5]string{"", "es", "", "", "secreto"}
var sato_msg = [5]string{"este", "", "un", "", ""}
var kenobi_msg_6 = [6]string{"este", "", "", "mensaje", "", ""}
var expectedMsg = "este es un mensaje secreto"*/

func TestValidateMessagesLen(t *testing.T) {
	if _, err := ValidateMessagesLen(kenobi_msg[:], skywalker_msg[:], sato_msg[:]); err != nil {
		t.Errorf("TestMsgLen FAILED: All messages were expected to have the same number of words.")
	} else {
		t.Log("TestMsgLen 1 PASSED.")
	}
	if valid, err := ValidateMessagesLen(kenobi_msg_6[:], skywalker_msg[:], sato_msg[:]); err == nil {
		t.Errorf("Expected: %v, got: %v", false, valid)
	} else {
		t.Log("TestMsgLen 2 PASSED.")
	}

}

func TestGetMessage(t *testing.T) {
	if msg := GetMessage(kenobi_msg[:], skywalker_msg[:], sato_msg[:]); msg != expectedMsg {
		t.Errorf("FAILED: Expected: %v, got: %v", expectedMsg, msg)
	} else {
		t.Log("TestExpectedMsg PASSED.")
	}
}

func TestGetJsonMessage(t *testing.T) {
	t.Log("TestGetJsonMessage...")
	var kenobi_msg = [5]string{"este", "", "", "mensaje", ""}
	var skywalker_msg = [5]string{"", "es", "", "", "secreto"}
	var sato_msg = [5]string{"este", "", "un", "", ""}
	resp := Response{Message: GetMessage(kenobi_msg[:], skywalker_msg[:], sato_msg[:])}
	body, err := json.Marshal(resp)
	if err != nil {
		t.Errorf("FAILED.. %v: ", err)
	}
	t.Log("body json: ", string(body))
	fmt.Println(string(body))
}
