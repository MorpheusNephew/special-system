package paperquotes

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/morpheusnephew/qotd/internal/utils"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	client httpClient
)

func init() {
	client = &http.Client{}
}

// GetQuoteOfTheDay gets the quote of the day and returns a QuoteOfTheDayResponse
func GetQuoteOfTheDay() (*QuoteOfTheDayResponse, *ErrorResponse) {
	qotdRequest := getQuoteOfTheDayRequest()

	return getQuoteOfTheDayResponse(qotdRequest)
}

func getGetRequest(url string, body io.Reader) *http.Request {
	return getRequest(http.MethodGet, url, body)
}

func getQuoteOfTheDayRequest() *http.Request {
	return getGetRequest("https://api.paperquotes.com/apiv1/qod/?lang=en", nil)
}

func getQuoteOfTheDayResponse(req *http.Request) (*QuoteOfTheDayResponse, *ErrorResponse) {
	body, errorResponse := getResponse(req)

	if errorResponse != nil {
		return nil, errorResponse
	}

	var r *QuoteOfTheDayResponse

	err := json.Unmarshal(body, &r)

	utils.PanicIfError(err)

	return r, nil
}

func getRequest(method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)

	utils.PanicIfError(err)

	authToken := os.Getenv("PAPER_QUOTES_TOKEN")

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", authToken))

	return req
}

func getResponse(req *http.Request) ([]byte, *ErrorResponse) {
	response, err := client.Do(req)

	utils.PanicIfError(err)

	defer response.Body.Close()

	var e *ErrorResponse

	if response.StatusCode >= 400 {
		e = &ErrorResponse{
			Code:    response.StatusCode,
			Message: response.Status[4:],
		}

		return nil, e
	}

	body, err := ioutil.ReadAll(response.Body)

	utils.PanicIfError(err)

	err = json.Unmarshal(body, &e)

	utils.PanicIfError(err)

	if e.Code > 0 {
		return nil, e
	}

	return body, nil
}
