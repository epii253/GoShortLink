package settings

import (
	"log"
	"errors"
	"os"
	"strconv"
	"sync"

	godotenv "github.com/joho/godotenv"
)

var once sync.Once

func LoadWithFallBack(path string) error {
	var err error

	once.Do(func() {
		if _, statErr := os.Stat(path); statErr == nil {
			err = godotenv.Load(path)
			return 
		}

		fallbackPath := path + ".example"

		if _, statErr := os.Stat(fallbackPath); statErr == nil {
			err = godotenv.Load(fallbackPath)
			return 
		}

		log.Println("no env files found; using system environment variables")


		err = errors.New("no env files found; using system environment variables")
	})

	return err
}

type Config struct {
	Host string
	Port int
}

func NewConfig(path string) *Config {
	LoadWithFallBack(path)

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	return &Config{Host: os.Getenv("HOST"), Port: port}
}
