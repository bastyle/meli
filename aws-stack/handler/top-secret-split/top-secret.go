package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := "Bastian"
	fmt.Printf("hello: %s", name)
	return events.APIGatewayProxyResponse{Body: "body proof ", StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
