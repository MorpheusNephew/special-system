package paperquotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/morpheusnephew/qotd/internal/testutils"
)

type SuccessfulMockClient struct {
}

func (c *SuccessfulMockClient) Do(req *http.Request) (*http.Response, error) {
	body := ioutil.NopCloser(strings.NewReader(`{ "data" : "hello world!!!"}`))
	return &http.Response{
		StatusCode: 200,
		Body:       body,
	}, nil
}

type UnsuccessfulMockClient struct {
}

func (c *UnsuccessfulMockClient) Do(req *http.Request) (*http.Response, error) {
	body := ioutil.NopCloser(strings.NewReader(`{ "data" : "hello world!!!"}`))
	return &http.Response{
		StatusCode: 400,
		Status:     "400 Bad Request",
		Body:       body,
	}, nil
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
	expectedToken := "myToken"
	os.Setenv("PAPER_QUOTES_TOKEN", expectedToken)

	r := getRequest(expectedMethod, expectedURL, nil)

	testutils.IfStringsNotEqual(t, r.Method, expectedMethod)
	testutils.IfStringsNotEqual(t, r.URL.String(), expectedURL)

	expectedContentType := "application/json"
	ct := r.Header.Get("Content-Type")
	testutils.IfStringsNotEqual(t, ct, expectedContentType)

	auth := r.Header.Get("Authorization")
	testutils.IfStringsNotEqual(t, auth, fmt.Sprintf("Token %v", expectedToken))
}

func Test_getResponse_Successfully(t *testing.T) {
	client = &SuccessfulMockClient{}

	body, _ := getResponse(&http.Request{})

	var res map[string]string

	json.Unmarshal(body, &res)

	testutils.IfStringsNotEqual(t, res["data"], "hello world!!!")
}

func Test_getResponse_Unsuccessfully(t *testing.T) {
	client = &UnsuccessfulMockClient{}

	_, err := getResponse(&http.Request{})

	testutils.IfIntsNotEqual(t, err.Code, 400)
	testutils.IfStringsNotEqual(t, err.Message, "Bad Request")
}
