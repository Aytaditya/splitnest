package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `yaml:"address"`
}

type Config struct {
	Env         string     `yaml:"env"`
	StoragePath string     `yaml:"storage_path"`
	HttpServer  HttpServer `yaml:"http_server"`
}

func LoadConfig() (*Config, error) {
	var configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flg := flag.String("config", "", "Path to configuration file")
		flag.Parse()
		configPath = *flg
		if configPath == "" {
			log.Fatal("Config Path is required")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("Config file does not exist at path:", configPath)
		return nil, err
	}
	// here we will read the file and unmarshal into config struct
	var cfg Config
	err1 := cleanenv.ReadConfig(configPath, &cfg)
	if err1 != nil {
		log.Fatal("Cannot read config file:", err1)
		return nil, err1
	}
	return &cfg, nil
}
