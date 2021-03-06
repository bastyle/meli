package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

//TODO crear funcion generica para retornar en casos de 500
type ExcpResponse struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

// function that receives the event triggered in aws.
func HandlePostOffLineRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("HandleOffLineRequest body = %v.\n", req.Body)

	satEntity := SatEntity{}
	err := json.Unmarshal([]byte(req.Body), &satEntity)
	if err != nil || satEntity.Distance == 0 || satEntity.Message == nil {
		fmt.Printf("not enough information: %v.\n", err.Error())
		//return events.APIGatewayProxyResponse{Body: string(`{"detail":"not enough information!}`), StatusCode: 500}, nil
		return events.APIGatewayProxyResponse{Body: getExceptionResponse("not enough information", 500), StatusCode: 500}, nil
	}
	// preparación de respuesta (transformacion)
	satelliteName := strings.ToLower(req.PathParameters["satellite_name"])
	fmt.Printf("satelliteName = %v.\n", satelliteName)
	if err := UpdateSatellite(satelliteName, satEntity.Distance, satEntity.Message); err != nil {
		return events.APIGatewayProxyResponse{Body: getExceptionResponse(err.Error(), 500), StatusCode: 500}, nil
	} else {
		return events.APIGatewayProxyResponse{Body: string(`{"detail":"satellite updated!}`), StatusCode: 200}, nil
	}
}

func HandleGetOffLineRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("HandleGetOffLineRequest body = %v.\n", req.Body)
	//obetener todos los satellites
	satellites, err := GetAllDataSatell()
	if err != nil {
		return events.APIGatewayProxyResponse{Body: getExceptionResponse(err.Error(), 500), StatusCode: 500}, nil
	} else if isOk, reqBody := IsThereNecessaryInfo(satellites); !isOk {
		reqBody.Satellites[0].Name = "solo para pasar"
		return events.APIGatewayProxyResponse{Body: getExceptionResponse("not enough information", 200), StatusCode: 200}, nil
	}
	//enviar a procesar satellites

	return events.APIGatewayProxyResponse{Body: getExceptionResponse("not enough information!", 500), StatusCode: 500}, nil
}

//function tah validates the information of the satllites in the db
func IsThereNecessaryInfo(satellites []SatEntity) (bool, RequestBody) {
	//TODO tengo que armar un objeto del tipo RequestBody y enviarlo a la funcion ProcessRequest de la version online
	reqBodyStruct := RequestBody{}
	//satArray := []Satellites{}

	/*if len(satellites) != satellitesExpected {
		fmt.Printf("satellites expected %v \n", satellitesExpected)
		return false, reqBodyStruct
	}*/

	//lenAux := 0
	fmt.Printf("reqBodyStruct:::: %v \n", reqBodyStruct)
	for _, s := range satellites {
		fmt.Printf("satellite %v \n", s)
		if s.Distance == 0 || len(s.Message) == 0 {
			fmt.Printf("not enough information. (distance: %v message: %v)\n", s.Distance, s.Message)
			return false, reqBodyStruct
		}
		fmt.Printf("dos::::::::::::::::::::::::::::::: %v \n", s)
		/*if lenAux == 0 { //primera iteracion
			lenAux = len(s.Message)
		}
		if lenAux != len(s.Message) {
			fmt.Printf("There is a difference with the length of the messages %v vs %v \n", lenAux, len(s.Message))
			return false, reqBodyStruct
		}*/
		//armar objeto por iteracion
		//reqBodyStruct = append(reqBodyStruct, s)
		//FIXME fuera de rango, hay que instanciar el arreglo
		satAux := Sats{}
		satAux.Name = s.Name
		fmt.Printf("tewa::::::::::::::::::::::::::::::: %v \n", s)
		satAux.Distance = float32(s.Distance)
		satAux.Message = s.Message
		reqBodyStruct.Satellites = append(reqBodyStruct.Satellites, satAux)
	}
	return true, reqBodyStruct
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
