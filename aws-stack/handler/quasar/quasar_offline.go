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
		return events.APIGatewayProxyResponse{Body: getExcptionResponse("not enough information!", 500), StatusCode: 500}, nil
	}
	// preparación de respuesta (transformacion)
	satelliteName := strings.ToLower(req.PathParameters["satellite_name"])
	fmt.Printf("satelliteName = %v.\n", satelliteName)
	if err := UpdateSatellite(satelliteName, satEntity.Distance, satEntity.Message); err != nil {
		return events.APIGatewayProxyResponse{Body: getExcptionResponse(err.Error(), 500), StatusCode: 500}, nil
	} else {
		return events.APIGatewayProxyResponse{Body: string(`{"detail":"satellite updated!}`), StatusCode: 200}, nil
	}
}

func getExcptionResponse(inputErr string, code int) string {
	excpRes := ExcpResponse{}
	excpRes.Detail = inputErr
	excpRes.Code = code
	if jsonExcBody, err := json.Marshal(excpRes); err != nil {
		return inputErr
	} else {
		return string(jsonExcBody)
	}
}
