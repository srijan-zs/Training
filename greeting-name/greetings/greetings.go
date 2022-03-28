package greetings

import "fmt"

func SayHello(name string) string {
	message := fmt.Sprintf("Hello %s!", name)
	return message
}
