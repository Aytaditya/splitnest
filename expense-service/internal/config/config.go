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
	Env         string     `yaml:"env"`
	Storagepath string     `yaml:"storage_path"`
	HttpServer  HttpServer `yaml:"http_server"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flg := flag.String("config", "", "Path to configuration file") // flat.string means we are expecting a string value from command line argument
		flag.Parse()
		configPath = *flg
		if configPath == "" {
			panic("Config Path is required")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("Config file does not exist at path: " + configPath)
	}
	var cfg Config
	err1 := cleanenv.ReadConfig(configPath, &cfg)
	if err1 != nil {
		panic("Cannot read config file: " + err1.Error())
	}

	return &cfg
}
