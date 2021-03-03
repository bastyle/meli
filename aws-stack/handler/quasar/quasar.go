package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("request.Body = %v.\n", req.Body)
	fmt.Printf("request.Body size = %d.\n", len(req.Body))
	fmt.Printf("path = %v.\n", req.Path)
	fmt.Printf("method = %v.\n", req.HTTPMethod)

	// request reading
	//reqBodyStruct := new(RequestBody)
	reqBodyStruct := RequestBody{}
	err := json.Unmarshal([]byte(req.Body), &reqBodyStruct)
	if err != nil {
		fmt.Printf("Error Unmarshal req.Body= %v.\n", err)
		return events.APIGatewayProxyResponse{Body: "Unable to unmarshal JSON req.body ", StatusCode: 500}, nil
	}
	// preparación de respuesta (transformacion)
	if jsonStrResp, err := ProcessRequest(reqBodyStruct); err != nil {
		fmt.Printf("Error ProcessRequest = %v.\n", err)
		//return events.APIGatewayProxyResponse{Body: "Error procesando mensaje.", StatusCode: 500}, nil
		return events.APIGatewayProxyResponse{StatusCode: 404}, nil
	} else {
		fmt.Printf("jsonStrResp: %v.\n", jsonStrResp)
		// response
		if jsonResBody, err := json.Marshal(jsonStrResp); err != nil {
			fmt.Printf("Error marshal jsonStrResp= %v.\n", err)
			//return events.APIGatewayProxyResponse{Body: "Error transformando response body a objeto JSON", StatusCode: 500}, nil
			return events.APIGatewayProxyResponse{StatusCode: 404}, nil
		} else {
			fmt.Printf("jsonStrResp %v: .\n", jsonStrResp)
			return events.APIGatewayProxyResponse{Body: string(jsonResBody), StatusCode: 200}, nil
		}
	}
}

//funcion encargada de procesar la petición
func ProcessRequest(reqBodyStruct RequestBody) (ResponseBody, error) {
	resBodyStruct := ResponseBody{}
	if len(reqBodyStruct.Satellites) != 3 {
		return resBodyStruct, errors.New("3 satellites are expected")
	} else {
		if msgResp, err := GetSecretMessage(reqBodyStruct); err != nil {
			return resBodyStruct, err
		} else {
			resBodyStruct.Message = msgResp
			return resBodyStruct, nil
		}
		//agregar llamada a obtener position
	}
}

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
		fmt.Printf("x: %v  y: %v", x, y)
		return x, y, nil
	}
}

func GetSecretMessage(reqBodyStruct RequestBody) (string, error) {
	fmt.Printf("msg_1: %v.\n", reqBodyStruct.Satellites[0].Message[:])
	if _, err := ValidateMessagesLen(reqBodyStruct.Satellites[0].Message[:], reqBodyStruct.Satellites[1].Message[:], reqBodyStruct.Satellites[2].Message[:]); err != nil {
		//error en largo de mensajes
		fmt.Printf("error:::: %v.\n", err)
		return "", err
	} else {
		fmt.Printf("validacion 1 ok %v.\n", err)
		resp := ResponseBody{Message: GetMessage(reqBodyStruct.Satellites[0].Message[:], reqBodyStruct.Satellites[1].Message[:], reqBodyStruct.Satellites[2].Message[:])}
		fmt.Printf("message %v.\n", resp.Message)
		return resp.Message, nil
	}

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
