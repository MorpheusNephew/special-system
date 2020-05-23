package paperquotes

import (
	"net/http"
	"testing"

	"github.com/morpheusnephew/qotd/internal/testutils"
)

func Test_getRequest(t *testing.T) {
	expectedMethod := http.MethodGet
	expectedURL := "www.testurl.com"

	r := getRequest(expectedMethod, expectedURL, nil)

	testutils.IfStringsNotEqual(t, r.Method, expectedMethod)
	testutils.IfStringsNotEqual(t, r.URL.String(), expectedURL)
}
