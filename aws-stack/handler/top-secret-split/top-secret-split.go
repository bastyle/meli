package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var topsecretSplitPath = "/topsecret_split"
var topsecretPath = "/topsecret"

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//name := "Bastian"
	//fmt.Printf("hello: %s.\n", name)
	//greet, _ := GetMessages(name)
	//fmt.Printf("req = %v.\n", req)
	//fmt.Printf("Path = %v.\n", req.Path)
	//pathAux := strings.TrimSuffix(req.Path, "/")
	pathAux := strings.Split(req.Path, "/")[1]
	fmt.Printf("pathAux = %v.\n", pathAux)
	fmt.Printf("method = %v.\n", req.HTTPMethod)
	satelliteName := strings.ToLower(req.PathParameters["satellite_name"])
	fmt.Printf("satelliteName = %v.\n", satelliteName)

	switch pathAux {
	case topsecretPath:
		fmt.Printf("redirect request to topsecret online api = %v.\n", satelliteName)
		return events.APIGatewayProxyResponse{Body: "redirect request to topsecret online api", StatusCode: 200}, nil
	case topsecretSplitPath:
		if "GET" == req.HTTPMethod {
			return events.APIGatewayProxyResponse{Body: "redirect request to get topsecret_split offline api" + satelliteName, StatusCode: 200}, nil
		} else if "POST" == req.HTTPMethod && "" != satelliteName {
			return events.APIGatewayProxyResponse{Body: "redirect request to post topsecret_split offline api" + satelliteName, StatusCode: 200}, nil
		} else {
			return events.APIGatewayProxyResponse{StatusCode: 404}, nil
		}
	default:
		return events.APIGatewayProxyResponse{StatusCode: 404}, nil
	}

	/*if topsecretPath == pathAux {
		fmt.Printf("redirect request to topsecret online api = %v.\n", satelliteName)
		return events.APIGatewayProxyResponse{Body: "redirect request to topsecret online api", StatusCode: 200}, nil
	} else if topsecretSplitPath == pathAux && "POST" == req.HTTPMethod && "" != satelliteName {
		return events.APIGatewayProxyResponse{Body: "topsecret_split post" + satelliteName, StatusCode: 200}, nil
	} else if topsecretSplitPath == pathAux && "GET" == req.HTTPMethod {
		return events.APIGatewayProxyResponse{Body: "topsecret_split get", StatusCode: 200}, nil
	}*/
	fmt.Printf("router = %v.\n", req)
	return events.APIGatewayProxyResponse{StatusCode: 404}, nil
}

func main() {
	lambda.Start(handleRequest)
}
