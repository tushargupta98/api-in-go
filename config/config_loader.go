package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var cfg *Config

func init() {
	cfg = loadConfig()
}

func loadConfig() *Config {
	var cfg *Config

	var envType = getEnv("RDM_ENVIRONMENT_TYPE")
	log.Info("Environemt Type: ", envType)
	var configPath = getConfigPath(envType)
	log.Info("Configle File Path: ", configPath)
	// Check if config file exists in current directory

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("Error loading env config file")
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal("Error unmarshaling env config file")
	}
	log.Info("Configuration loaded successfully from integration environment's directory")

	PopulateEnvVariables(cfg)
	return cfg
}

func UpdateEnvConfig() {
	cfg = nil
}

func GetConfig() *Config {
	return cfg
}

func PopulateEnvVariables(config *Config) {
	// Set DatabaseConfig values from environment variables
	config.Database.Host = getEnv(config.Database.Host)
	config.Database.Name = getEnv(config.Database.Name)
	config.Database.Schema = getEnv(config.Database.Schema)
	config.Database.User = getEnv(config.Database.User)
	config.Database.Password = getEnv(config.Database.Password)
	config.Database.SSLRootCert = getEnv(config.Database.SSLRootCert)

	// Set CacheConfig values from environment variables
	config.Cache.Host = getEnv(config.Cache.Host)
	config.Cache.DB = getEnv(config.Cache.DB)

	// Set NewRelicConfig values from environment variables
	config.NewRelic.LicenseKey = getEnv(config.NewRelic.LicenseKey)
	config.NewRelic.AppName = getEnv(config.NewRelic.AppName)
}

// Helper function to retrieve an environment variable.
// If the environment variable is not set, an error message is printed and the program exits.
func getEnv(key string) string {
	value := os.Getenv(key)
	return value
}

func getConfigPath(envType string) string {
	switch envType {
	case "local":
		return "config/env/config.local.yaml"
	case "int":
		return "/var/www/app/config/env/config.int.yaml"
	case "stag":
		return "/var/www/app/config/env/config.stag.yaml"
	default:
		return ""
	}
}
