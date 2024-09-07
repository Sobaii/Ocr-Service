package config

import "os"

type Config struct {
	OpenAIAPIKey string
	Port         string
}

func Load() *Config {
	return &Config{
		OpenAIAPIKey: os.Getenv("OPENAI_API_KEY"),
		Port:         os.Getenv("PORT"),
	}
}