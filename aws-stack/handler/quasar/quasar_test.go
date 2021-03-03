package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

var kenobi_msg = [5]string{"este", "", "", "mensaje", ""}
var skywalker_msg = [5]string{"", "es", "", "", "secreto"}
var sato_msg = [5]string{"este", "", "un", "", ""}
var kenobi_msg_6 = [6]string{"este", "", "", "mensaje", "", ""}
var expectedMsg = "este es un mensaje secreto"

func TestExampleMessage(t *testing.T) {
	if msg := GetExampleMessage(); msg.Message != expectedMsg {
		t.Errorf("FAILED: Expected: %v, got: %v", expectedMsg, msg)
	} else {
		t.Log("TestExpectedMsg PASSED")
	}
}

func TestJsonRequest(t *testing.T) {
	t.Log("TestJsonRequest ...")
	//body := `{"satellites":[{"name":"kenobi","distance":100.0,"message":["este","","","mensaje",""]},{"name":"skywalker","distance":115.5,"message":["","es","","","secreto"]},{"name":"sato","distance":142.7,"message":["este","","un","",""]}]}`
	body := `"{"satellites":[{"name":"kenobi","distance":100.0,"message":["este","","","mensaje",""]},{"name":"skywalker","distance":115.5,"message":["","es","","","secreto"]},{"name":"sato","distance":142.7,"message":["este","","un","",""]}]}"`

	if last := len(body) - 1; last >= 0 && body[last] == '"' {
		body = body[:last]
	}
	if len(body) >= 0 && body[0] == '"' {
		body = body[1:]
	}

	t.Log("TestJsonRequest body: ", body)
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Errorf("FAILED: %v", err)
	} else {
		t.Log("TestJsonRequest PASSED. body: ", jsonBody, "err: ", err)
	}
	//bodyJSON := make(map[string][]RequestBody)
	bodyJSON := RequestBody{}
	err1 := json.Unmarshal([]byte(body), &bodyJSON)
	if err1 != nil {
		panic(err1)
	}
	//fmt.Printf("\n\n json object:::: %+v", bodyJSON)
	//fmt.Printf("\n\n json object:::: %+v", bodyJSON.satellites[0].Name)
	fmt.Printf("\n\n json object:::: %+v", bodyJSON.Satellites[0].Name)
}

func TestByteStr(t *testing.T) {
	b := []byte{34, 65, 66, 67, 226, 130, 172, 34}
	fmt.Printf("b: %v \n", b)
	s := string(b)
	if b[0] == 34 {
		fmt.Printf("quotes: %v \n", b[0])
		b = b[1 : len(b)-1]
	}
	fmt.Printf("b: %v \n", b)
	fmt.Printf("str: %v \n", s)
}

/*func TrimSuffix(s, suffix string) string {
    if strings.HasSuffix(s, suffix) {
        s = s[:len(s)-len(suffix)]
    }
    return s
}*/
