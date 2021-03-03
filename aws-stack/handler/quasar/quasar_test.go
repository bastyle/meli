package main

import (
	"encoding/json"
	"testing"
)

var kenobi_msg = [5]string{"este", "", "", "mensaje", ""}
var skywalker_msg = [5]string{"", "es", "", "", "secreto"}
var sato_msg = [5]string{"este", "", "un", "", ""}
var kenobi_msg_6 = [6]string{"este", "", "", "mensaje", "", ""}
var expectedMsg = "este es un mensaje secreto"
var bodyJsonExample = `{"satellites":[{"name":"kenobi","distance":100.0,"message":["este","","","mensaje",""]},{"name":"skywalker","distance":115.5,"message":["","es","","","secreto"]},{"name":"sato","distance":142.7,"message":["este","","un","",""]}]}`

func TestExampleMessage(t *testing.T) {
	if msg := GetExampleMessage(); msg.Message != expectedMsg {
		t.Errorf("FAILED: Expected: %v, got: %v", expectedMsg, msg)
	}
}

func TestGetSecretMessage(t *testing.T) {
	reqBodyStruct := RequestBody{}
	err := json.Unmarshal([]byte(bodyJsonExample), &reqBodyStruct)
	if err != nil {
		t.Errorf("Error Unmarshal req.Body= %v.\n", err)
	} else {
		//t.Log("reqBodyStruct.Satellites[0].Name :", reqBodyStruct.Satellites[0].Name)
		if jsonStrResp, err := GetSecretMessage(reqBodyStruct); err != nil {
			t.Errorf("Error getting satellites message= %v.\n", err)
		} else {
			t.Log("jsonStrResp: ", jsonStrResp)
		}
	}
}

func TestProcessRequest(t *testing.T) {
	reqBodyStruct := RequestBody{}
	err := json.Unmarshal([]byte(bodyJsonExample), &reqBodyStruct)
	if err != nil {
		t.Errorf("Error Unmarshal req.Body= %v.\n", err)
	} else {
		if resBodyStruct, err := ProcessRequest(reqBodyStruct); err != nil {
			t.Errorf("Error ProcessRequest = %v.\n", err)
		} else {
			t.Log("resBodyStruct: ", resBodyStruct)
		}
	}
}

func TestGetPosition(t *testing.T) {
	reqBodyStruct := RequestBody{}
	err := json.Unmarshal([]byte(bodyJsonExample), &reqBodyStruct)
	if err != nil {
		t.Errorf("Error Unmarshal req.Body= %v.\n", err)
	} else {
		if x, y, err := GetPosition(reqBodyStruct); err != nil {
			t.Errorf("Error ProcessRequest = %v.\n", err)
		} else {
			t.Log("x: ", x, "y: ", y)
		}
	}
}
