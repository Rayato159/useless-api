package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type IConfig interface {
	App() IAppConfig
}

type config struct {
	app *app
}

func NewConfig(path string) IConfig {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &config{
		app: &app{
			url: os.Getenv("APP_URL"),
		},
	}
}

type IAppConfig interface {
	Url() string
}

type app struct {
	url string
}

func (cfg *config) App() IAppConfig { return cfg.app }
func (a *app) Url() string          { return a.url }
