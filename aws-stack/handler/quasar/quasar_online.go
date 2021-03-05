package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

type MockResponse struct {
	Message string `json:"message"`
}

type RequestBody struct {
	Satellites []struct {
		Name     string   `json:"name"`
		Distance float32  `json:"distance"`
		Message  []string `json:"message"`
	} `json:"satellites"`
}

type ResponseBody struct {
	Position struct {
		X float32 `json:"x"`
		Y float32 `json:"y"`
	} `json:"position"`
	Message string `json:"message"`
}

// function that receives the event triggered in aws.
func handleOffLineRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("request.Body = %v.\n", req.Body)
	fmt.Printf("request.Body size = %d.\n", len(req.Body))
	fmt.Printf("path = %v.\n", req.Path)
	fmt.Printf("method = %v.\n", req.HTTPMethod)
	reqBodyStruct := RequestBody{}
	err := json.Unmarshal([]byte(req.Body), &reqBodyStruct)
	if err != nil {
		fmt.Printf("Error Unmarshal req.Body= %v.\n", err)
		return events.APIGatewayProxyResponse{Body: "Unable to unmarshal JSON req.body ", StatusCode: 500}, nil
	}
	// preparación de respuesta (transformacion)
	if responseBody, err := ProcessRequest(reqBodyStruct); err != nil {
		fmt.Printf("Error ProcessRequest = %v.\n", err)
		//return events.APIGatewayProxyResponse{Body: "Error procesando mensaje.", StatusCode: 500}, nil
		return events.APIGatewayProxyResponse{StatusCode: 404}, nil
	} else { // response
		fmt.Printf("responseBody: %v.\n", responseBody)
		if jsonResBody, err := json.Marshal(responseBody); err != nil {
			fmt.Printf("Error marshal responseBody= %v.\n", err)
			return events.APIGatewayProxyResponse{StatusCode: 404}, nil
		} else {
			fmt.Printf("jsonResBody %v: .\n", jsonResBody)
			return events.APIGatewayProxyResponse{Body: string(jsonResBody), StatusCode: 200}, nil
		}
	}
}

//this function is in charge of process the request.
func ProcessRequest(reqBodyStruct RequestBody) (ResponseBody, error) {
	resBodyStruct := ResponseBody{}
	if len(reqBodyStruct.Satellites) != 3 {
		return resBodyStruct, errors.New("3 satellites are expected")
	} else {
		if secretMsg, err := GetSecretMessage(reqBodyStruct); err != nil {
			return ResponseBody{}, err
		} else {
			resBodyStruct.Message = secretMsg
		}
		if x, y, err := GetPosition(reqBodyStruct); err != nil {
			return ResponseBody{}, err
		} else {
			resBodyStruct.Position.X = x
			resBodyStruct.Position.Y = y
			return resBodyStruct, nil
		}
	}
}

//this function is a wrapper to get position.
func GetPosition(reqBodyStruct RequestBody) (float32, float32, error) {
	kenobiDistance := reqBodyStruct.Satellites[0].Distance
	skywalkerDistance := reqBodyStruct.Satellites[1].Distance
	satoDistance := reqBodyStruct.Satellites[2].Distance
	// chequear tipo de dato mejor
	if kenobiDistance == 0 || skywalkerDistance == 0 || satoDistance == 0 {
		return 0, 0, errors.New("Distances must be greater than 0.")
	}
	if x, y := GetLocation(kenobiDistance, skywalkerDistance, satoDistance); &x == nil || &y == nil {
		return x, y, errors.New("Error calculating position")
	} else {
		fmt.Printf("x: %v  y: %v.\n", x, y)
		return x, y, nil
	}
}

//this function is a wrapper to get secret message.
func GetSecretMessage(reqBodyStruct RequestBody) (string, error) {
	//fmt.Printf("msg_1: %v.\n", reqBodyStruct.Satellites[0].Message[:])
	if _, err := ValidateMessagesLen(reqBodyStruct.Satellites[0].Message[:], reqBodyStruct.Satellites[1].Message[:], reqBodyStruct.Satellites[2].Message[:]); err != nil {
		fmt.Printf("error:::: %v.\n", err)
		return "", err
	} else {
		resp := ResponseBody{Message: GetMessage(reqBodyStruct.Satellites[0].Message[:], reqBodyStruct.Satellites[1].Message[:], reqBodyStruct.Satellites[2].Message[:])}
		fmt.Printf("secret message obtained: %v.\n", resp.Message)
		return resp.Message, nil
	}

}

// the main function
/*func main() {
	lambda.Start(handleRequest)
}*/
