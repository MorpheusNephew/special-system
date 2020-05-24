package paperquotes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

func Test_getRequest(t *testing.T) {
	expectedMethod := http.MethodGet
	expectedURL := "www.testurl.com"

	r := getRequest(expectedMethod, expectedURL, nil)

	testutils.IfStringsNotEqual(t, r.Method, expectedMethod)
	testutils.IfStringsNotEqual(t, r.URL.String(), expectedURL)
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
