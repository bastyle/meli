package messages

import (
	"errors"
	"strings"
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
			if finalMsg[i] == "" {
				finalMsg[i] = v[i]
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
