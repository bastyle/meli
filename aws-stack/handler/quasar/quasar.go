package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Message struct {
	Name string
	Body string
	Age  int
}

type Response struct {
	Message string `json:"message"`
}

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(getMessage())
	//body, err := json.Marshal(getMessage2())
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Unable to marshal JSON", StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func getMessage() Message {
	return Message{"Bastian", "Hello", 37}
}

/*func getMessage2() Response {
	var kenobi_msg = [5]string{"este", "", "", "mensaje", ""}
	var skywalker_msg = [5]string{"", "es", "", "", "secreto"}
	var sato_msg = [5]string{"este", "", "un", "", ""}
	return Response{Message: GetMessage(kenobi_msg[:], skywalker_msg[:], sato_msg[:])}
}*/

func main() {
	lambda.Start(handleRequest)
}
