package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var topsecretSplitPath = "topsecret_split"
var topsecretPath = "topsecret"

const satellitesExpected = 3

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pathAux := strings.Split(req.Path, "/")[1]
	fmt.Printf("::::pathAux = %v.\n", pathAux)
	fmt.Printf("method = %v.\n", req.HTTPMethod)
	satelliteName := strings.ToLower(req.PathParameters["satellite_name"])
	fmt.Printf("satelliteName = %v.\n", satelliteName)

	switch pathAux {
	case topsecretPath:
		fmt.Printf("redirecting request to topsecret online api = %v.\n", satelliteName)
		//return events.APIGatewayProxyResponse{Body: "redirect request to topsecret online api", StatusCode: 200}, nil
		return HandleOnLineRequest(req)
	case topsecretSplitPath:
		if "GET" == req.HTTPMethod {
			return events.APIGatewayProxyResponse{Body: "redirect request to get topsecret_split offline api" + satelliteName, StatusCode: 200}, nil
		} else if "POST" == req.HTTPMethod && "" != satelliteName {
			//return events.APIGatewayProxyResponse{Body: "redirect request to post topsecret_split offline api" + satelliteName, StatusCode: 200}, nil
			fmt.Printf("redirecting request to post topsecret_split offline api = %v.\n", satelliteName)
			return HandlePostOffLineRequest(req)
		} else {
			return events.APIGatewayProxyResponse{StatusCode: 404}, nil
		}
	default:
		return events.APIGatewayProxyResponse{StatusCode: 404}, nil
	}
	fmt.Printf("req = %v.\n", req)
	return events.APIGatewayProxyResponse{StatusCode: 404}, nil
}

func main() {
	lambda.Start(handleRequest)
}
