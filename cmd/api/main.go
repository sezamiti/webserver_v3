package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/sezamiti/go2/StandardWebServer/internal/app/api"
	"log"
	"os"
)

var (
	configPath   string
	configFormat string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path ti config file in .toml format")
	flag.StringVar(&configFormat, "format", "toml", "config file format (toml or env)")
}
func main() {
	flag.Parse()
	log.Println("it works")

	config := api.NewConfig()

	if configPath != "" {
		switch configFormat {
		case "toml":
			if _, err := toml.DecodeFile(configPath, config); err != nil {
				log.Println("Cannot find or decode TOML config file. Using default values.", err)
			}
		case "env":
			if err := loadEnvConfig(configPath, config); err != nil {
				log.Println("Cannot load env config file. Using default values.", err)
			}
		default:
			log.Println("Invalid config format. Using default values.")
		}
		server := api.New(config)

		if err := server.Start(); err != nil {
			log.Fatal(err)
		}
	}

}

func loadEnvConfig(path string, config *api.Config) error {
	envVars, err := readEnvFile(path)
	if err != nil {
		return err
	}

	if bindAddr, ok := envVars["BIND_ADDR"]; ok {
		config.BindAddr = bindAddr
	}

	if loggerLevel, ok := envVars["LOGGER_LEVEL"]; ok {
		config.LoggerLevel = loggerLevel
	}

	// Загрузите остальные переменные окружения, если необходимо

	return nil
}

func readEnvFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open env file: %w", err)
	}
	defer file.Close()

	envVars, err := godotenv.Parse(file)
	if err != nil {
		return nil, fmt.Errorf("failed to parse env file: %w", err)
	}

	return envVars, nil
}
