package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var topsecretSplitPath = "/topsecret_split"

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := "Bastian"
	fmt.Printf("hello: %s.\n", name)
	//greet, _ := GetMessages(name)
	fmt.Printf("req = %v.\n", req)
	fmt.Printf("Path = %v.\n", req.Path)
	pathAux := strings.TrimSuffix(req.Path, "/")
	fmt.Printf("pathAux = %v.\n", pathAux)
	fmt.Printf("method = %v.\n", req.HTTPMethod)
	satelliteName := req.PathParameters["satellite_name"]
	fmt.Printf("satellite_name pathparam = %v.\n", req.PathParameters["satellite_name"])
	fmt.Printf("satelliteName = %v.\n", satelliteName)
	if "/topsecret" == pathAux {
		return events.APIGatewayProxyResponse{Body: "topsecret", StatusCode: 200}, nil
	} else if topsecretSplitPath == pathAux && "POST" == req.HTTPMethod && "" != satelliteName {
		return events.APIGatewayProxyResponse{Body: "topsecret_split post" + satelliteName, StatusCode: 200}, nil
	} else if topsecretSplitPath == pathAux && "GET" == req.HTTPMethod {
		return events.APIGatewayProxyResponse{Body: "topsecret_split get", StatusCode: 200}, nil
	}
	fmt.Printf("router = %v.\n", req)
	return events.APIGatewayProxyResponse{Body: name, StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
