package variables

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	// RedisKey is the key used for getting/setting quote of the day in Redis
	RedisKey string

	// PaperQuotesToken is the token used to get the quote of the day
	PaperQuotesToken string
)

func init() {
	godotenv.Load()

	RedisKey = os.Getenv("REDIS_KEY")
	PaperQuotesToken = os.Getenv("PAPER_QUOTES_TOKEN")
}
