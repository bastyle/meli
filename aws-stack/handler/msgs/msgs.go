package msgs

import "fmt"

/*type MockResp struct {
	Message string `json:"message"`
}*/

func GetMessages(name string) (string, error) {
	fmt.Print("hello....")
	return "hello " + name, nil
}
