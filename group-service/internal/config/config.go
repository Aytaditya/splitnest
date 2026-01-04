package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `yaml:"address"`
}

type Config struct {
	ENV         string     `yaml:"env"`
	StoragePath string     `yaml:"storage_path"`
	HttpServer  HttpServer `yaml:"http_server"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flg := flag.String("config", "", "Path to configuration file")
		flag.Parse()
		configPath = *flg
		if configPath == "" {
			panic("config path is required")
		}
	}

	if _, err := os.Stat(configPath); err != nil {
		panic("failed to access config file: " + err.Error())
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config file: " + err.Error())
	}

	return &cfg
}
