package messages

import (
	"fmt"
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
func GetMessage(messages []string) string {
	fmt.Println("GetMessage....", messages)
	//fmt.Println(messages)
	//message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return messages[0]
}

func GetMessage2(messages ...[]string) string {
	fmt.Println("GetMessage22....", messages)
	fmt.Println("GetMessage22 arreglos: ", len(messages))
	//fmt.Println(messages)
	//message := fmt.Sprintf("Hi, %v. Welcome!", name)
	//finalMsg := []string
	var finalMsg []string

	for counter, v := range messages {
		fmt.Println("counter: ", counter)
		if counter == 0 { //primera iteración
			finalMsg = v[:]
			fmt.Println("contador 0, igualando arreglo")
		}
		fmt.Println("finalMsg len: ", len(finalMsg))
		//fmt.Println("finalMsg: ", finalMsg)
		/*if 0 == len(finalMsg) {
			fmt.Println("msg final vacio, se inicializa...")
			finalMsg := v
			fmt.Println("----", finalMsg)
			fmt.Println("finalMsg22 len: ", len(finalMsg))
		}*/

		fmt.Println("arreglo de entrada: ", v, "largo: ", len(v))
		//fmt.Println("largo: ", len(v))
		//fmt.Println(v, "--", reflect.ValueOf(v).Kind())
		//for i, m := range v {
		for i := 0; i < len(v); i++ {
			//fmt.Println("final msg[: ", i, "] :", finalMsg[i])
			fmt.Println("valor: ", v[i])
			if finalMsg[i] == "" {
				finalMsg[i] = v[i]
			}
		}
		fmt.Println("finalMsg: ", finalMsg)
	}
	//return messages[0][0]
	return finalMsg[0]
}
