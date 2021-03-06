package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
		initializeAPI()
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

func initializeAPI() {
	fmt.Println("API things")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET"},
	}))

	router.GET("/qotd", func(c *gin.Context) {

		response, errorResponse := paperquotes.GetQuoteOfTheDay()

		if errorResponse != nil {
			c.JSON(errorResponse.Code, gin.H{
				"error": gin.H{
					"code":    errorResponse.Code,
					"message": errorResponse.Message,
				},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": response,
			})
		}
	})

	router.Run(":3333")
}
