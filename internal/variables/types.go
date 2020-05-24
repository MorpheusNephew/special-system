package variables

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	// PaperQuotesToken is the token used to get the quote of the day
	PaperQuotesToken string

	// RedisKey is the key used for getting/setting quote of the day in Redis
	RedisKey string
)

func init() {
	godotenv.Load()

	PaperQuotesToken = os.Getenv("PAPER_QUOTES_TOKEN")
	RedisKey = os.Getenv("REDIS_KEY")
}
