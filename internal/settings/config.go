package settings

import (
	"log"
	"os"
	"strconv"
	"sync"

	godotenv "github.com/joho/godotenv"
)

var once sync.Once

func LoadWithFallBack(path string) {
	once.Do(func() {
		if err := godotenv.Load(path); err != nil {

			// fallback
			if err := godotenv.Load(path + ".example"); err != nil {
				log.Println("no env files found, relying on system env")
			}
		}
	})
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
