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

type Response struct {
	Message string `json:"message"`
}

type RequestBody struct {
	Satellites []struct {
		Name     string   `json:"name"`
		Distance float64  `json:"distance"`
		Message  []string `json:"message"`
	} `json:"satellites"`
}

type ResponseBody struct {
	Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"position"`
	Message string `json:"message"`
}

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
	// request reading
	//reqBodyStruct := new(RequestBody)
	reqBodyStruct := RequestBody{}
	err := json.Unmarshal([]byte(req.Body), reqBodyStruct)
	if err != nil {
		fmt.Printf("Error Unmarshal req.Body= %v.\n", err)
		return events.APIGatewayProxyResponse{Body: "Unable to unmarshal JSON req.body ", StatusCode: 500}, nil
	} /*else {
		fmt.Printf("reqBodyStruct.Satellites[0].Name = %v.\n", reqBodyStruct.Satellites[0].Name)
		//resp := ResponseBody{Message: GetSecretMessage(reqBodyStruct)}
	}*/
	// preparación de respuesta (transformacion)

	if jsonStrResp, err := ProcessRequest(reqBodyStruct); err != nil {
		fmt.Printf("Error ProcessRequest = %v.\n", err)
		return events.APIGatewayProxyResponse{Body: "Error procesando mensaje.", StatusCode: 500}, nil
	} else {
		fmt.Printf("jsonStrResp: %v.\n", jsonStrResp)
		// response
		if jsonResBody, err := json.Marshal(jsonStrResp); err != nil {
			fmt.Printf("Error marshal jsonStrResp= %v.\n", err)
			return events.APIGatewayProxyResponse{Body: "Error transformando response body a objeto JSON", StatusCode: 500}, nil
		} else {
			fmt.Printf("jsonStrResp %v: .\n", jsonStrResp)
			return events.APIGatewayProxyResponse{Body: string(jsonResBody), StatusCode: 200}, nil
		}
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

//funcion encargada de procesar la petición
func ProcessRequest(reqBodyStruct RequestBody) (ResponseBody, error) {
	resBodyStruct := ResponseBody{}
	if msgResp, err := GetSecretMessage(reqBodyStruct); err != nil {
		return resBodyStruct, err
	} else {
		resBodyStruct.Message = msgResp
		return resBodyStruct, nil
	}
}

func GetSecretMessage(reqBodyStruct RequestBody) (string, error) {
	//validar que el objeto tenga 3 satellites
	if len(reqBodyStruct.Satellites) != 3 {
		//fmt.Errorf("")
		return "", errors.New("3 satellites are expected")
	} else {
		//crear un arreglo de arreglos para enviar
		//var satellitesMessages [3][]string{}
		fmt.Printf("msg_1: %v.\n", reqBodyStruct.Satellites[0].Message[:])
		if _, err := ValidateMessagesLen(reqBodyStruct.Satellites[0].Message[:], reqBodyStruct.Satellites[1].Message[:], reqBodyStruct.Satellites[2].Message[:]); err != nil {
			//error en largo de mensajes
			fmt.Printf("error:::: %v", err)
			return "", err
		} else {
			fmt.Printf("validacion 1 ok %v.\n", err)
			resp := ResponseBody{Message: GetMessage(reqBodyStruct.Satellites[0].Message[:], reqBodyStruct.Satellites[1].Message[:], reqBodyStruct.Satellites[2].Message[:])}

			fmt.Printf("message %v.\n", resp.Message)
			return resp.Message, nil
		}
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
