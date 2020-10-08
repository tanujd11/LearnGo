package main

import (
	"fmt"

	greetings "github.com/tanujd11/LearnGo"

	"log"
)

func main() {
	message, error := greetings.Hello("")

	if error == nil {
		fmt.Println(message)
	} else {
		log.Fatal(error)
	}
}
