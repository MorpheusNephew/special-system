package paperquotes

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/morpheusnephew/qotd/internal/redisclient"
	"github.com/morpheusnephew/qotd/internal/utils"
	"github.com/morpheusnephew/qotd/internal/variables"
)

type iHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	client             iHTTPClient
	redisClientFactory redisclient.IClientFactory
)

func init() {
	client = &http.Client{}
	redisClientFactory = &redisclient.ClientFactory{}
}

// GetQuoteOfTheDay gets the quote of the day and returns a QuoteOfTheDayResponse
func GetQuoteOfTheDay() (*QuoteOfTheDayResponse, *ErrorResponse) {
	redisKey := fmt.Sprintf("%v-qotd", variables.RedisKeyPrefix)

	redisClientFactory.GetRedisClient().GetValue(redisKey)

	body, errorResponse := retrieveData(redisKey, func() ([]byte, *ErrorResponse) {
		qotdRequest := getQuoteOfTheDayRequest()

		return getResponse(redisKey, qotdRequest)
	})

	if errorResponse != nil {
		return nil, errorResponse
	}

	quoteOfTheDayResponse, errorResponse := getQuoteOfTheDayResponse(body)

	return quoteOfTheDayResponse, errorResponse
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

func getResponse(redisKey string, req *http.Request) ([]byte, *ErrorResponse) {
	response, err := client.Do(req)

	utils.PanicIfError(err)

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return nil, &ErrorResponse{
			Code:    response.StatusCode,
			Message: response.Status[4:],
		}
	}

	body, err := ioutil.ReadAll(response.Body)

	utils.PanicIfError(err)

	var cacheTTL *time.Duration = nil
	expiresHeader := response.Header.Get("Expires")

	if len(expiresHeader) > 0 {
		expiresHeaderGMT, err := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", expiresHeader)

		utils.PanicIfError(err)

		// TODO (JJ): Find a way to get the absolute value of the tme difference
		*cacheTTL = expiresHeaderGMT.UTC().Sub(time.Now().UTC())
	}

	redisClientFactory.GetRedisClient().SetValue(redisKey, body, cacheTTL)

	return body, nil
}

func retrieveData(redisKey string, f func() ([]byte, *ErrorResponse)) ([]byte, *ErrorResponse) {
	cacheResponse, _ := redisClientFactory.GetRedisClient().GetValue(redisKey)

	if len(cacheResponse) > 0 {
		return cacheResponse, nil
	}

	body, errorResponse := f()

	return body, errorResponse
}
