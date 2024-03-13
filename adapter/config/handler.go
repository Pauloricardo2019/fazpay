package config

import (
	"fmt"
	"github.com/Pauloricardo2019/teste_fazpay/internal/model"
	"github.com/joho/godotenv"
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
			Environment: os.Getenv("ENV"),
			BasePath:    os.Getenv("BASE_PATH"),
		}

		config.DBConfig = model.DBConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
		}

		fmt.Println("Conn string: ", config.DBConfig.GetConnString())

		restPort := os.Getenv("PORT")

		intRestPort, err := strconv.Atoi(restPort)
		if err != nil {
			intRestPort = 8080
		}

		config.RestPort = intRestPort
	})

	return config
}
