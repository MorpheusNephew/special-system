package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/morpheusnephew/qotd/internal/paperquotes"
	"github.com/morpheusnephew/qotd/internal/variables"
)

// HandleRequest is the logic that will be ran when a lambda has been invoked
func HandleRequest(ctx context.Context) (*paperquotes.QuoteOfTheDayResponse, error) {
	return getQuoteOfTheDay()
}

func main() {
	env := strings.ToLower(variables.Environment)

	switch env {

	case "lambda":
		lambda.Start(HandleRequest)
		break

	case "api":
		fmt.Println("API things")
		break

	default:
		getQuoteOfTheDay()
	}
}

func getQuoteOfTheDay() (*paperquotes.QuoteOfTheDayResponse, error) {
	fmt.Println("Starting to get Quote of the Day")

	var err error
	response, errorResponse := paperquotes.GetQuoteOfTheDay()

	if errorResponse != nil {
		err = fmt.Errorf("%v %v", errorResponse.Code, errorResponse.Message)
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}

	return response, err
}
