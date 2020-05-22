package main

import (
	"fmt"
	"log"

	"github.com/morpheusnephew/qotd/internal/paperquotes"
)

func main() {

	response, errorResponse := paperquotes.GetQuoteOfTheDay()

	if errorResponse != nil {
		log.Fatalln(errorResponse)
	} else {
		fmt.Println(response)
	}
}
