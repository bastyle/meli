package main

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// sattelitesMsgs messages struct
type sattelitesMsgs struct {
	kenobi    []string
	skywalker []string
	sato      []string
}

// GetMessage get a complete message from an array of messages
func GetMessage(messages ...[]string) string {
	var finalMsg []string
	for counter, v := range messages {
		if counter == 0 { //primera iteración
			finalMsg = v[:]
		}
		for i := 0; i < len(v); i++ {
			if strings.TrimSpace(finalMsg[i]) == "" {
				finalMsg[i] = strings.TrimSpace(v[i])
			}
		}
	}
	return strings.Join(finalMsg[:], " ")
}

// ValidateMessagesLen validates that all arrays have the same size
func ValidateMessagesLen(messages ...[]string) (bool, error) {
	var msgLen int
	msgLen = len(messages[0])
	for i := 1; i < len(messages); i++ {
		if msgLen != len(messages[i]) {
			return false, errors.New("la cantidad de palabras debe ser igual en todos los mensajes.")
		}
	}
	return true, nil
}

/********aws*******/
type Message struct {
	Name string
	Body string
	age  int
}

type Response struct {
	Message string `json:"message"`
}

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//body, err := json.Marshal(getMessage())
	body, err := json.Marshal(getMessage2())
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Unable to marshal JSON", StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func getMessage() Message {
	return Message{"Alice", "Hello", 37}
}

func getMessage2() Response {
	var kenobi_msg = [5]string{"este", "", "", "mensaje", ""}
	var skywalker_msg = [5]string{"", "es", "", "", "secreto"}
	var sato_msg = [5]string{"este", "", "un", "", ""}
	return Response{Message: GetMessage(kenobi_msg[:], skywalker_msg[:], sato_msg[:])}
//	return Response{GetMessage(kenobi_msg[:], skywalker_msg[:], sato_msg[:])}
}

func main() {
	lambda.Start(handleRequest)
}
