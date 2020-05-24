package variables

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	// PaperQuotesToken is the token used to get the quote of the day
	PaperQuotesToken string

	// RedisEndpoint is the endpoint for interacting with Redis
	RedisEndpoint string

	// RedisKey is the key used for getting/setting quote of the day in Redis
	RedisKey string

	// RedisPort is the port used with interacting with Redis
	RedisPort string
)

func init() {
	godotenv.Load()

	PaperQuotesToken = os.Getenv("PAPER_QUOTES_TOKEN")
	RedisEndpoint = os.Getenv("REDIS_ENDPOINT")
	RedisKey = os.Getenv("REDIS_KEY")
	RedisPort = os.Getenv("REDIS_PORT")
}
