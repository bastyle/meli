package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

const notEnoughInfor = "not enough information"

type ExcpResponse struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

// function that receives the post req triggered in aws.
func HandlePostOffLineRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("HandleOffLineRequest body = %v.\n", req.Body)
	satEntity := SatEntity{}
	err := json.Unmarshal([]byte(req.Body), &satEntity)
	if err != nil || satEntity.Distance == 0 || satEntity.Message == nil {
		fmt.Printf("not enough information: %v.\n", err.Error())
		//return events.APIGatewayProxyResponse{Body: string(`{"detail":"not enough information!}`), StatusCode: 500}, nil
		return events.APIGatewayProxyResponse{Body: getExceptionResponse(notEnoughInfor, 500), StatusCode: 500}, nil
	}
	// preparación de respuesta (transformacion)
	satelliteName := strings.ToLower(req.PathParameters["satellite_name"])
	fmt.Printf("satelliteName = %v.\n", satelliteName)
	if err := UpdateSatellite(satelliteName, satEntity.Distance, satEntity.Message); err != nil {
		return events.APIGatewayProxyResponse{Body: getExceptionResponse(err.Error(), 500), StatusCode: 500}, nil
	} else {
		return events.APIGatewayProxyResponse{Body: string(`{"detail":"satellite updated!"}`), StatusCode: 200}, nil
	}
}

// function that receives the get req triggered in aws.
func HandleGetOffLineRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("HandleGetOffLineRequest body = %v.\n", req.Body)
	//obetener todos los satellites
	satellites, err := GetAllDataSatell()
	if err != nil {
		return events.APIGatewayProxyResponse{Body: getExceptionResponse(err.Error(), 500), StatusCode: 500}, nil
	}
	if isOk, reqBody := IsThereNecessaryInfo(satellites); !isOk {
		return events.APIGatewayProxyResponse{Body: getExceptionResponse(notEnoughInfor, 200), StatusCode: 200}, nil
	} else if responseBody, err := ProcessRequest(reqBody); err != nil {
		fmt.Printf("Error ProcessRequest = %v.\n", err)
		return events.APIGatewayProxyResponse{Body: getExceptionResponse(notEnoughInfor, 200), StatusCode: 200}, nil
	} else {
		fmt.Printf("responseBody: %v.\n", responseBody)
		if jsonResBody, err := json.Marshal(responseBody); err != nil {
			fmt.Printf("Error marshal responseBody= %v.\n", err)
			return events.APIGatewayProxyResponse{Body: getExceptionResponse(notEnoughInfor, 200), StatusCode: 200}, nil
		} else {
			//fmt.Printf("jsonResBody %v: .\n", jsonResBody)
			return events.APIGatewayProxyResponse{Body: string(jsonResBody), StatusCode: 200}, nil
		}
	}
	return events.APIGatewayProxyResponse{Body: getExceptionResponse(notEnoughInfor, 500), StatusCode: 500}, nil
}

//function tah validates the information of the satllites in the db
func IsThereNecessaryInfo(satellites []SatEntity) (bool, RequestBody) {
	//TODO tengo que armar un objeto del tipo RequestBody y enviarlo a la funcion ProcessRequest de la version online
	reqBodyStruct := RequestBody{}
	fmt.Printf("reqBodyStruct: %v \n", reqBodyStruct)
	for _, s := range satellites {
		fmt.Printf("satellite %v \n", s)
		if s.Distance == 0 || len(s.Message) == 0 {
			fmt.Printf("not enough information. (distance: %v message: %v)\n", s.Distance, s.Message)
			return false, reqBodyStruct
		}
	}
	//ordenar arreglo (kenobi, skywalker, sato) darle un id al satelite? y ordenar por ahí
	var sats [3]Sats
	sats[0] = getSatByName("kenobi", satellites)
	sats[1] = getSatByName("skywalker", satellites)
	sats[2] = getSatByName("sato", satellites)
	reqBodyStruct.Satellites = sats[:]
	return true, reqBodyStruct
}

// func that get satellite by name
func getSatByName(satName string, satellites []SatEntity) Sats {
	for _, s := range satellites {
		if satName == s.Name {
			satAux := Sats{}
			satAux.Name = s.Name
			satAux.Distance = float32(s.Distance)
			satAux.Message = s.Message
			return satAux
		}
	}
	return Sats{}
}

//function that generates a generic json error
func getExceptionResponse(inputErr string, code int) string {
	excpRes := ExcpResponse{}
	excpRes.Detail = inputErr
	excpRes.Code = code
	if jsonExcBody, err := json.Marshal(excpRes); err != nil {
		return inputErr
	} else {
		return string(jsonExcBody)
	}
}
