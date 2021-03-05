package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := "Bastian"
	fmt.Printf("hello: %s.\n", name)
	//greet, _ := GetMessages(name)
	fmt.Printf("req = %v.\n", req)
	fmt.Printf("Path = %v.\n", req.Path)
	fmt.Printf("method = %v.\n", req.HTTPMethod)
	fmt.Printf("QueryStringParameters = %v.\n", req.QueryStringParameters["satellite_name"])
	fmt.Printf("pathparam = %v.\n", req.PathParameters["satellite_name"])
	if "/topsecret" == req.Path {
		return events.APIGatewayProxyResponse{Body: "topsecret", StatusCode: 200}, nil
	} else if "/topsecret_split" == req.Path && "POST" == req.HTTPMethod && "" != req.QueryStringParameters["satellite_name"] {
		return events.APIGatewayProxyResponse{Body: "topsecret_split post" + req.QueryStringParameters["satellite_name"], StatusCode: 200}, nil
	} else if "/topsecret_split" == req.Path && "GET" == req.HTTPMethod {
		return events.APIGatewayProxyResponse{Body: "topsecret_split get", StatusCode: 200}, nil
	}
	fmt.Printf("router = %v.\n", req)
	return events.APIGatewayProxyResponse{Body: name, StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
