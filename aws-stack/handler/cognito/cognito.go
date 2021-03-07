package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("body:: %v \n", req.Body)
	return events.APIGatewayProxyResponse{Body: "OK", StatusCode: 200}, nil
}

// the main function
func main() {
	lambda.Start(handleRequest)
}
