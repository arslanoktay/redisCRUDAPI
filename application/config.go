package application

import (
	"os"
	"strconv"
)

type Config struct {
	RedisAdress string
	Serverport  uint16
}

// We're not gonna use viper because to undestand what's happening
func LoadConfig() Config {
	cfg := Config{
		RedisAdress: "localhost:6379",
		Serverport:  3000,
	}

	if redisAddr, exists := os.LookupEnv("REDIS_ADDR"); exists {
		cfg.RedisAdress = redisAddr
	}

	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			cfg.Serverport = uint16(port)
		}
	}

	return cfg
}
