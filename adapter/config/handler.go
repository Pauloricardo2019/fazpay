package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"kickoff/internal/model"
	"os"
	"strconv"
	"sync"
)

var config *model.Config
var doOnce sync.Once

type GetConfigFn func() *model.Config

func GetConfig() *model.Config {
	doOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
		}

		config = &model.Config{
			Environment:  os.Getenv("ENV"),
			DbConnString: os.Getenv("DB_CONNSTRING"),
			BasePath:     os.Getenv("BASE_PATH"),
		}

		restPort := os.Getenv("PORT")

		intRestPort, err := strconv.Atoi(restPort)
		if err != nil {
			intRestPort = 8080
		}

		config.RestPort = intRestPort
	})

	return config
}
