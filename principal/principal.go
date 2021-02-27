package main

import (
	"fmt"

	"meli.com/messages"
)

func main() {
	var arrayint = [...]int{1, 2, 3, 4, 5} //assign
	fmt.Println(arrayint)
	/*println("Array of arrays:")
	var arrayofarrays [3][len(arrayint)]int
	for i := range arrayofarrays { //assign
		arrayofarrays[i] = arrayint
	}
	fmt.Println(arrayofarrays, "arrayofarrays")*/

	fmt.Println("::::::::::::::::::::")
	//message := messages.Hello("Gladys")
	//fmt.Println(message)

	kenobi_msg := [5]string{"este", "", "", "mensaje", ""}
	fmt.Println("kenobi_msg 1:: ", kenobi_msg[:])
	skywalker_msg := [5]string{"", "es", "", "", "secreto"}
	sato_msg := [5]string{"este", "", "un", "", ""}
	arrayMessages := [][]string{kenobi_msg[:], skywalker_msg[:], sato_msg[:]}
	fmt.Println("arrayMessages::::", arrayMessages)

	//messages.GetMessage(kenobi_msg[:])
	messages.GetMessage2(kenobi_msg[:], skywalker_msg[:], sato_msg[:])
	//messages.GetMessage2(arrayMessages[:])
}

func ValidateMessagesLen() bool {
	return false
}
