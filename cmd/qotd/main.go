package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/morpheusnephew/qotd/internal/paperquotes"
	"github.com/morpheusnephew/qotd/internal/variables"
)

// HandleRequest is the logic that will be ran when a lambda has been invoked
func HandleRequest(ctx context.Context) (*paperquotes.QuoteOfTheDayResponse, *paperquotes.ErrorResponse) {
	return getQuoteOfTheDay()
}

func main() {
	if strings.ToLower(variables.Environment) == "local" {
		getQuoteOfTheDay()
	} else {
		lambda.Start(HandleRequest)
	}
}

func getQuoteOfTheDay() (*paperquotes.QuoteOfTheDayResponse, *paperquotes.ErrorResponse) {
	response, errorResponse := paperquotes.GetQuoteOfTheDay()

	if errorResponse != nil {
		err := fmt.Errorf("%v %v", errorResponse.Code, errorResponse.Message)
		log.Fatalln(err)
	} else {
		fmt.Println(response)
	}

	return response, errorResponse
}
