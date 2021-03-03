package main

import (
	"encoding/json"
	"fmt"

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

/*func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(GetExampleMessage())
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Unable to marshal JSON", StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}*/

/*func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	jsonBody, err := json.Marshal(req)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Unable to marshal JSON" + string(err.), StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(jsonBody), StatusCode: 200}, nil
}*/

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("request.Body = %v.\n", req.Body)
	fmt.Printf("request.Body size = %d.\n", len(req.Body))
	//reqBodyStruct := RequestBody{}
	reqBodyStruct := new(RequestBody)
	err := json.Unmarshal([]byte(req.Body), reqBodyStruct)
	//acá
	if err != nil {
		fmt.Printf("Error unMarshal req.Body= %v.\n", err)
		return events.APIGatewayProxyResponse{Body: "Unable to unmarshal JSON req.body ", StatusCode: 500}, nil
	} else {
		fmt.Printf("reqBodyStruct.Satellites[0].Name = %v.\n", reqBodyStruct.Satellites[0].Name)
		reqBodyStruct.Satellites[0].Name = "Prueba"
		fmt.Printf("reqBodyStruct.Satellites[0].Name = %v.\n", reqBodyStruct.Satellites[0].Name)
	}
	if jsonResBody, err := json.Marshal(reqBodyStruct); err != nil {
		fmt.Printf("Error marshal reqBodyStruct= %v.\n", err)
		return events.APIGatewayProxyResponse{Body: "Error transformando response body a objeto JSON", StatusCode: 500}, nil
	} else {
		fmt.Printf("transformacion de response ok %v: .\n", err)
		return events.APIGatewayProxyResponse{Body: string(jsonResBody), StatusCode: 200}, nil
	}
}

/*func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))
	//agregar trim
	jsonBody, err := json.Marshal(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Unable to marshal JSON", StatusCode: 500}, nil
	} else {
		fmt.Printf("reqJsonBody: %v.\n", string(jsonBody))
	}
	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}
	if last := len(string(jsonBody)) - 1; last >= 0 && string(jsonBody)[last] == '"' {
		jsonBody = jsonBody[:last]
	}
	if len(body) >= 0 && body[0] == '"' {
		body = body[1:]
	}
	reqBodyStruct := RequestBody{}
	err1 := json.Unmarshal([]byte(jsonBody), &reqBodyStruct)
	if err1 != nil {
		fmt.Printf("err: %v .\n", err1)
		return events.APIGatewayProxyResponse{Body: "Error transformando request body a JSON" + string(jsonBody), StatusCode: 500}, nil
	}
	reqBodyStruct.Satellites[0].Name = "Prueba"
	if jsonResBody, err := json.Marshal(reqBodyStruct); err != nil {
		return events.APIGatewayProxyResponse{Body: "Error transformando body a objeto JSON", StatusCode: 500}, nil
	} else {
		return events.APIGatewayProxyResponse{Body: string(jsonResBody), StatusCode: 200}, nil
	}
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}*/

func GetExampleMessage() MockResponse {
	var kenobi_msg = [5]string{"este", "", "", "mensaje", ""}
	var skywalker_msg = [5]string{"", "es", "", "", "secreto"}
	var sato_msg = [5]string{"este", "", "un", "", ""}
	return MockResponse{Message: GetMessage(kenobi_msg[:], skywalker_msg[:], sato_msg[:])}
}

func main() {
	lambda.Start(handleRequest)
}
