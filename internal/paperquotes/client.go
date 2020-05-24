package paperquotes

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/morpheusnephew/qotd/internal/redisclient"
	"github.com/morpheusnephew/qotd/internal/utils"
	"github.com/morpheusnephew/qotd/internal/variables"
)

type iHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	client      iHTTPClient
	redisClient redisclient.IClient
)

func init() {
	client = &http.Client{}
	redisClient = redisclient.GetRedisClient()
}

// GetQuoteOfTheDay gets the quote of the day and returns a QuoteOfTheDayResponse
func GetQuoteOfTheDay() (*QuoteOfTheDayResponse, *ErrorResponse) {
	redisKey := fmt.Sprintf("%v-qotd", variables.RedisKeyPrefix)

	redisClient.GetValue(redisKey)

	body, errorResponse := retrieveData(redisKey, func() ([]byte, *ErrorResponse) {
		qotdRequest := getQuoteOfTheDayRequest()

		return getResponse(qotdRequest)
	})

	if errorResponse != nil {
		return nil, errorResponse
	}

	quoteOfTheDayResponse, errorResponse := getQuoteOfTheDayResponse(body)

	return quoteOfTheDayResponse, errorResponse
}

func retrieveData(redisKey string, f func() ([]byte, *ErrorResponse)) ([]byte, *ErrorResponse) {
	cacheResponse, _ := redisClient.GetValue(redisKey)

	if len(cacheResponse) > 0 {
		return cacheResponse, nil
	}

	body, errorResponse := f()

	redisClient.SetValue(redisKey, body)

	return body, errorResponse
}

func getGetRequest(url string, body io.Reader) *http.Request {
	return getRequest(http.MethodGet, url, body)
}

func getQuoteOfTheDayRequest() *http.Request {
	return getGetRequest("https://api.paperquotes.com/apiv1/qod/?lang=en", nil)
}

func getQuoteOfTheDayResponse(body []byte) (*QuoteOfTheDayResponse, *ErrorResponse) {
	var r *QuoteOfTheDayResponse

	err := json.Unmarshal(body, &r)

	utils.PanicIfError(err)

	return r, nil
}

func getRequest(method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)

	utils.PanicIfError(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", variables.PaperQuotesToken))

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
