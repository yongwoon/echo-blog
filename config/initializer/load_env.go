package initializer

import (
	"os"

	"github.com/joho/godotenv"
)

func InitDotenv() {
	currentEnvironment, ok := os.LookupEnv("ENVIRONMENT")

	var err error
	if ok == true {
		err = godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/yongwoon/echo-blog-api/.env." + currentEnvironment))
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		panic(err)
	}
}
