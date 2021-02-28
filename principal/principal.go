package main

import (
	"fmt"
	"os"

	"meli.com/messages"
)

/*func ValidateMessagesLenLocal(messages ...[]string) (bool, error) {
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
}*/

func main() {
	var arrayint = [...]int{1, 2, 3, 4, 5} //assign
	fmt.Println(arrayint)

	fmt.Println(":::::PRINCIPAL:::::::::::")
	//message := messages.Hello("Gladys")
	//fmt.Println(message)
	//datos de entrada
	kenobi_msg := [5]string{"este", "", "", "mensaje", ""}
	//fmt.Println("kenobi_msg 1:: ", kenobi_msg[:])
	skywalker_msg := [5]string{"", "es", "", "", "secreto"}
	//sato_msg := [5]string{"este", "", "un", "", ""}
	sato_msg := [6]string{"este", "", "un", "", "", ""}
	arrayMessages := [][]string{kenobi_msg[:], skywalker_msg[:], sato_msg[:]}
	fmt.Println("arrayMessages::::", arrayMessages)

	//messages.GetMessage(kenobi_msg[:])
	//var finalMsg string = messages.GetMessage2(kenobi_msg[:], skywalker_msg[:], sato_msg[:])
	//fmt.Println("finalMsg::: ", finalMsg)
	/*if messages.ValidateMessagesLen(kenobi_msg[:], skywalker_msg[:], sato_msg[:]) {
		var finalMsg string = messages.GetMessage2(kenobi_msg[:], skywalker_msg[:], sato_msg[:])
		fmt.Println("finalMsg::: ", finalMsg)
	} */

	if _, err := messages.ValidateMessagesLen(kenobi_msg[:], skywalker_msg[:], sato_msg[:]); err != nil {
		fmt.Println("Err::: ", err)
	} else {
		var finalMsg string = messages.GetMessage(kenobi_msg[:], skywalker_msg[:], sato_msg[:])
		fmt.Println("finalMsg::: ", finalMsg)
		os.Exit(3)
	}
}
