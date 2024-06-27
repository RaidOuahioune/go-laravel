package job

import (
	"github.com/joho/godotenv"

	"os"
)

func GetRedisAddress() string {
	godotenv.Load(".env")

	return os.Getenv("REDIS_ADDR")
}
