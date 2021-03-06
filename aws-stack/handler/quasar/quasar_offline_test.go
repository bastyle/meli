package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestIsThereNecessaryInfo(t *testing.T) {
	if items, err := GetAllDataSatell(); err != nil {
		t.Errorf("Error GetAllDataSatell = %v.\n", err)
	} else {
		if isOk, requestBody := IsThereNecessaryInfo(items); !isOk {
			t.Errorf("Error IsThereNecessaryInfo = %v.\n", isOk)
		} else {
			t.Log("requestBody: ", requestBody)
		}
	}
}

func TestGetExceptionResponse(t *testing.T) {
	if m := getExceptionResponse("example error", 404); m == "" {
		t.Errorf("excp err example = %v.\n", m)
	}
}

func TestOnliGetHandler(t *testing.T) {
	req := events.APIGatewayProxyRequest{Body: ""}
	_, err := HandleGetOffLineRequest(req)
	if err != nil {
		t.Errorf("err: %v\n", err.Error())
	}
	//t.Log("TestGetHandler resp::::::: ", res)
}

func TestOfflineHandlePostOffLineRequest(t *testing.T) {
	var bodyAux = `{"distance":538.51,"message":["este","","","mensaje",""]}`
	var pathParams = map[string]string{"satellite_name": "sato"}
	req := events.APIGatewayProxyRequest{Body: bodyAux, PathParameters: pathParams}
	_, err := HandlePostOffLineRequest(req)
	if err != nil {
		t.Errorf("err: %v\n", err.Error())
	}
	//t.Log("TestGetHandler resp::::::: ", res)
}
