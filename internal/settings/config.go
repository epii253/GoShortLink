package settings

import (
	"errors"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

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

	DBHost string
	DBPort int
	DBUser string
	DBPass string
	DBName string

	RedisAddr    string
	RedisPass    string
	CacheTTL     time.Duration
}

func NewConfig(path string) *Config {
	LoadWithFallBack(path)

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return &Config{
		Host: os.Getenv("HOST"),
		Port: port,

		RedisAddr: os.Getenv("REDIS_ADDR"),
		RedisPass: os.Getenv("REDIS_PASSWORD"),
		CacheTTL: func() time.Duration {
			if d, err := time.ParseDuration(os.Getenv("CACHE_TTL")); err == nil {
				return d
			}
			return 2 * time.Minute
		}(),

		DBHost: os.Getenv("DB_HOST"),
		DBPort: dbPort,
		DBUser: os.Getenv("POSTGRES_USER"),
		DBPass: os.Getenv("POSTGRES_PASSWORD"),
		DBName: os.Getenv("POSTGRES_DB"),
	}
}