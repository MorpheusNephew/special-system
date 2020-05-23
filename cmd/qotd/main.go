package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/morpheusnephew/qotd/internal/paperquotes"
)

func main() {
	godotenv.Load()

	response, errorResponse := paperquotes.GetQuoteOfTheDay()

	if errorResponse != nil {
		log.Fatalln(errorResponse)
	} else {
		fmt.Println(response)
	}
}
