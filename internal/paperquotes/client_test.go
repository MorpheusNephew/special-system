package paperquotes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/morpheusnephew/qotd/internal/testutils"
)

type MockClient struct {
}

func (c *MockClient) Do(req *http.Request) (*http.Response, error) {
	body := ioutil.NopCloser(strings.NewReader(`{ "data" : "hello world!!!"}`))
	return &http.Response{
		StatusCode: 200,
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

func Test_getResponse(t *testing.T) {
	client = &MockClient{}

	body, err := getResponse(&http.Request{})

	if err != nil {
		t.Errorf("Expected %v, but received %v", nil, err)
	}

	var res map[string]string

	json.Unmarshal(body, &res)

	testutils.IfStringsNotEqual(t, res["data"], "hello world!!!")
}
