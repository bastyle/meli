package messages

import (
	"errors"
	"fmt"
	"strings"
)

const PartsOfMessage = 5

type sattelitesMsgs struct {
	kenobi    []string
	skywalker []string
	sato      []string
}

func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	fmt.Println(fmt.Sprintf("%d %s", PartsOfMessage, "parts: "))
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

//func GetMessage(messages [PartsOfMessage]string) string {
func GetMessageIni(messages []string) string {
	fmt.Println("GetMessage....", messages)
	//fmt.Println(messages)
	//message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return messages[0]
}

func GetMessage(messages ...[]string) string {
	//fmt.Println(":::: GetMessage2 ::::", messages, "cantidad de arreglos: ", len(messages))
	var finalMsg []string
	for counter, v := range messages {
		//fmt.Println("counter: ", counter)
		if counter == 0 { //primera iteración
			finalMsg = v[:]
			//fmt.Println("contador 0, igualando arreglo")
		}
		//fmt.Println("finalMsg len: ", len(finalMsg))
		//fmt.Println("finalMsg: ", finalMsg)

		//fmt.Println("arreglo de entrada: ", v, "largo: ", len(v))
		//fmt.Println("largo: ", len(v))
		//fmt.Println(v, "--", reflect.ValueOf(v).Kind())
		//for i, m := range v {
		for i := 0; i < len(v); i++ {
			//fmt.Println("final msg[: ", i, "] :", finalMsg[i])
			//fmt.Println("valor: ", v[i])
			if finalMsg[i] == "" {
				finalMsg[i] = v[i]
			}
		}
		//fmt.Println("finalMsg: ", finalMsg)
	}
	//return messages[0][0]
	//return finalMsg[0]
	return strings.Join(finalMsg[:], " ")
}

func ValidateMessagesLen(messages ...[]string) (bool, error) {
	var msgLen int
	msgLen = len(messages[0])
	//for counter, v := range messages {
	for i := 1; i < len(messages); i++ {
		if msgLen != len(messages[i]) {
			//fmt.Println("existe un arreglo con un largo diferente: ")
			return false, errors.New("la cantidad de palabras debe ser igual en todos los mensajes.")
		}
	}
	return true, nil
}
