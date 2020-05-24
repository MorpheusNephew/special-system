package main

import (
	"fmt"
	"log"

	"github.com/morpheusnephew/qotd/internal/paperquotes"
)

func main() {
	response, errorResponse := paperquotes.GetQuoteOfTheDay()

	if errorResponse != nil {
		err := fmt.Errorf("%v %v", errorResponse.Code, errorResponse.Message)
		log.Fatalln(err)
	} else {
		fmt.Println(response)
	}
}
