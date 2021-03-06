package main

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

var kenobi_msg = [5]string{"este", "", "", "mensaje", ""}
var skywalker_msg = [5]string{"", "es", "", "", "secreto"}
var sato_msg = [5]string{"este", "", "un", "", ""}
var kenobi_msg_6 = [6]string{"este", "", "", "mensaje", "", ""}
var expectedMsg = "este es un mensaje secreto"
var bodyJsonExample = `{"satellites":[{"name":"kenobi","distance":100.0,"message":["este","","","mensaje",""]},{"name":"skywalker","distance":115.5,"message":["","es","","","secreto"]},{"name":"sato","distance":142.7,"message":["este","","un","",""]}]}`

func TestGetSecretMessage(t *testing.T) {
	reqBodyStruct := RequestBody{}
	err := json.Unmarshal([]byte(bodyJsonExample), &reqBodyStruct)
	if err != nil {
		t.Errorf("Error Unmarshal req.Body= %v.\n", err)
	} else {
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

func TestHandleOnLineRequest(t *testing.T) {
	var bodyAux = `{"satellites":[{"name":"kenobi","distance":100.0,"message":["este","","","mensaje",""]},{"name":"skywalker","distance":115.5,"message":["","es","","","secreto"]},{"name":"sato","distance":142.7,"message":["este","","un","",""]}]}`
	//var pathParams = map[string]string{"satellite_name": "sato"}
	req := events.APIGatewayProxyRequest{Body: bodyAux}
	res, err := HandleOnLineRequest(req)
	if err != nil {
		t.Errorf("err: %v\n", err.Error())
	} else {
		t.Log("TestGetHandler status code: ", res.StatusCode)
	}

}
