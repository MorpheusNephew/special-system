package paperquotes

// QuoteOfTheDayResponse is the response data from the request
// GET https://api.paperquotes.com/apiv1/qod/
type QuoteOfTheDayResponse struct {
	Author   string   `json:"author"`
	Language string   `json:"language"`
	Likes    int      `json:"likes"`
	Quote    string   `json:"quote"`
	Tags     []string `json:"tags"`
}
