package paperquotes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/morpheusnephew/qotd/internal/redisclient"
	"github.com/morpheusnephew/qotd/internal/testutils"
)

type SuccessfulMockClient struct {
}

type UnsuccessfulMockClient struct {
}

type MockRedisClient struct {
}

type MockRedisClientFactory struct {
	client redisclient.IClient
}

func (c *SuccessfulMockClient) Do(req *http.Request) (*http.Response, error) {
	body := ioutil.NopCloser(strings.NewReader(`{ "data" : "hello world!!!"}`))
	return &http.Response{
		StatusCode: 200,
		Body:       body,
	}, nil
}

func (c *UnsuccessfulMockClient) Do(req *http.Request) (*http.Response, error) {
	body := ioutil.NopCloser(strings.NewReader(`{ "data" : "hello world!!!"}`))
	return &http.Response{
		StatusCode: 400,
		Status:     "400 Bad Request",
		Body:       body,
	}, nil
}

func (c *MockRedisClient) GetValue(key string) ([]byte, error) {

	return []byte(`{ "data" : "hello world!!!"}`), nil
}

func (c *MockRedisClient) SetValue(key string, value []byte, t *time.Duration) (string, error) {

	return "Set", nil
}

func (c *MockRedisClient) GetInitialized() bool {
	return true
}

func (c *MockRedisClient) SetInitialized(value bool) {
}

func (cf *MockRedisClientFactory) GetRedisClient() redisclient.IClient {
	return &MockRedisClient{}
}

func Test_getQuoteOfTheDayResponse(t *testing.T) {
	expectedQuoteOfTheDayResponse := &QuoteOfTheDayResponse{
		Author:   "Confucius",
		Likes:    200,
		Language: "en",
		Tags:     []string{"Inspirational"},
		Quote:    "Those who say they can and those who say they can't are both right",
	}

	data, _ := json.Marshal(expectedQuoteOfTheDayResponse)

	quoteOfTheDayResponse, _ := getQuoteOfTheDayResponse(data)

	testutils.IfStringsNotEqual(t, quoteOfTheDayResponse.Quote, expectedQuoteOfTheDayResponse.Quote)
	testutils.IfStringsNotEqual(t, quoteOfTheDayResponse.Author, expectedQuoteOfTheDayResponse.Author)
	testutils.IfStringsNotEqual(t, quoteOfTheDayResponse.Language, expectedQuoteOfTheDayResponse.Language)
	testutils.IfStringsNotEqual(t, quoteOfTheDayResponse.Tags[0], expectedQuoteOfTheDayResponse.Tags[0])
	testutils.IfIntsNotEqual(t, quoteOfTheDayResponse.Likes, expectedQuoteOfTheDayResponse.Likes)
}

func Test_getRequest(t *testing.T) {
	expectedMethod := http.MethodGet
	expectedURL := "www.testurl.com"

	r := getRequest(expectedMethod, expectedURL, nil)

	testutils.IfStringsNotEqual(t, r.Method, expectedMethod)
	testutils.IfStringsNotEqual(t, r.URL.String(), expectedURL)

	expectedContentType := "application/json"
	ct := r.Header.Get("Content-Type")
	testutils.IfStringsNotEqual(t, ct, expectedContentType)

	auth := r.Header.Get("Authorization")
	testutils.IfStringIsEmpty(t, auth)
}

func Test_getResponse_Successfully(t *testing.T) {
	client = &SuccessfulMockClient{}
	redisClientFactory = &MockRedisClientFactory{}

	body, _ := getResponse("test", &http.Request{})

	var res map[string]string

	json.Unmarshal(body, &res)

	testutils.IfStringsNotEqual(t, res["data"], "hello world!!!")
}

func Test_getResponse_Unsuccessfully(t *testing.T) {
	client = &UnsuccessfulMockClient{}

	_, err := getResponse("test", &http.Request{})

	testutils.IfIntsNotEqual(t, err.Code, 400)
	testutils.IfStringsNotEqual(t, err.Message, "Bad Request")
}
