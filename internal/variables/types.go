package variables

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	// Environment is the environment that is currently running this applications
	Environment string

	// PaperQuotesToken is the token used to get the quote of the day
	PaperQuotesToken string

	// RedisEndpoint is the endpoint for interacting with Redis
	RedisEndpoint string

	// RedisKeyPrefix is the key used for getting/setting quote of the day in Redis
	RedisKeyPrefix string

	// RedisPort is the port used with interacting with Redis
	RedisPort string
)

func init() {
	godotenv.Load()

	Environment = os.Getenv("ENV")
	PaperQuotesToken = os.Getenv("PAPER_QUOTES_TOKEN")
	RedisEndpoint = os.Getenv("REDIS_ENDPOINT")
	RedisKeyPrefix = os.Getenv("REDIS_KEY_PREFIX")
	RedisPort = os.Getenv("REDIS_PORT")
}
