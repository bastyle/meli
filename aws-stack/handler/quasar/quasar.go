package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MockResponse struct {
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

/*type Satellite struct {
	Nombre   string   `json:"nombre"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type RequestBody struct {
	Satellites []Satellite `json:"satellites"`
}*/

type RequestBody struct {
	Satellites []struct {
		Name     string   `json:"name"`
		Distance float64  `json:"distance"`
		Message  []string `json:"message"`
	} `json:"satellites"`
}

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(GetExampleMessage())
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Unable to marshal JSON", StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func GetExampleMessage() MockResponse {
	var kenobi_msg = [5]string{"este", "", "", "mensaje", ""}
	var skywalker_msg = [5]string{"", "es", "", "", "secreto"}
	var sato_msg = [5]string{"este", "", "un", "", ""}
	return MockResponse{Message: GetMessage(kenobi_msg[:], skywalker_msg[:], sato_msg[:])}
}

func main() {
	lambda.Start(handleRequest)
}
