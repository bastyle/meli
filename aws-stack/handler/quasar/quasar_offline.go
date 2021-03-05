package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// function that receives the event triggered in aws.
func HandlePostOffLineRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("HandleOffLineRequest body = %v.\n", req.Body)

	satEntity := SatEntity{}
	err := json.Unmarshal([]byte(req.Body), &satEntity)
	if err != nil || satEntity.Distance == 0 || satEntity.Message == nil {
		fmt.Printf("not enough information: %v.\n", err.Error())
		return events.APIGatewayProxyResponse{Body: "not enough information", StatusCode: 500}, nil
	}
	// preparación de respuesta (transformacion)
	satelliteName := strings.ToLower(req.PathParameters["satellite_name"])
	fmt.Printf("satelliteName = %v.\n", satelliteName)
	if err := UpdateSatellite(satelliteName, satEntity.Distance, satEntity.Message); err != nil {
		return events.APIGatewayProxyResponse{Body: string(err.Error()), StatusCode: 500}, nil
	} else {
		return events.APIGatewayProxyResponse{StatusCode: 200}, nil
	}
}
