package main

import (
	"fmt"

	"github.com/morpheusnephew/qotd/internal/paperquotes"
)

func main() {

	response := &paperquotes.QuoteOfTheDayResponse{
		Author:   "Confucius",
		Quote:    "Those who say they can and those who say they can't are both usually right.",
		Language: "en",
		Likes:    10,
		Tags:     []string{"Inspirational"},
	}

	fmt.Println(response)
}
