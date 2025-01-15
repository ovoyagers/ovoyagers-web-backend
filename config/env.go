package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config is a struct that represents the configuration settings.
type Config struct {
	ENVIRONMENT                 string `mapstructure:"ENVIRONMENT"`
	PORT                        string `mapstructure:"PORT"`
	NEO4J_USER                  string `mapstructure:"NEO4J_USER"`
	NEO4J_PASSWORD              string `mapstructure:"NEO4J_PASSWORD"`
	NEO4J_URI                   string `mapstructure:"NEO4J_URI"`
	TWILLIO_ACCOUNT_SID         string `mapstructure:"TWILLIO_ACCOUNT_SID"`
	TWILLIO_ACCOUNT_PASSWORD    string `mapstructure:"TWILLIO_ACCOUNT_PASSWORD"`
	TWILLIO_ACCOUNT_SERVICE_SID string `mapstructure:"TWILLIO_ACCOUNT_SERVICE_SID"`
	ACCESS_TOKEN_SECRET         string `mapstructure:"ACCESS_TOKEN_SECRET"`
	REFRESH_TOKEN_SECRET        string `mapstructure:"REFRESH_TOKEN_SECRET"`
	EMAIL_HOST                  string `mapstructure:"EMAIL_HOST"`
	EMAIL_PORT                  string `mapstructure:"EMAIL_PORT"`
	EMAIL_USER                  string `mapstructure:"EMAIL_USER"`
	EMAIL_PASSWORD              string `mapstructure:"EMAIL_PASSWORD"`
	GIN_MODE                    string `mapstructure:"GIN_MODE"`
	IMAGEKIT_PUBLIC_KEY         string `mapstructure:"IMAGEKIT_PUBLIC_KEY"`
	IMAGEKIT_PRIVATE_KEY        string `mapstructure:"IMAGEKIT_PRIVATE_KEY"`
	IMAGEKIT_URL_ENDPOINT       string `mapstructure:"IMAGEKIT_URL_ENDPOINT"`
	UPSTASH_REDIS_REST_URL      string `mapstructure:"UPSTASH_REDIS_REST_URL"`
	UPSTASH_REDIS_REST_TOKEN    string `mapstructure:"UPSTASH_REDIS_REST_TOKEN"`
}

// LoadConfig loads the configuration from the specified path.
func LoadConfig() (c *Config, err error) {
	environment, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		environment = "local"
	}

	if environment == "local" {
		viper.SetConfigFile(".env")
		viper.AutomaticEnv() // Load matching environment variables automatically

		if err = viper.ReadInConfig(); err != nil {
			log.Error("Error reading configuration file: ", err)
			return nil, err
		}

		if err = viper.Unmarshal(&c); err != nil {
			log.Error("Error unmarshaling configuration: ", err)
			return nil, err
		}
	} else {
		// List of required environment variables
		envVars := []string{
			"ACCESS_TOKEN_SECRET", "REFRESH_TOKEN_SECRET", "PORT", "NEO4J_USER", "NEO4J_PASSWORD",
			"NEO4J_URI", "TWILLIO_ACCOUNT_SID", "TWILLIO_ACCOUNT_PASSWORD", "TWILLIO_ACCOUNT_SERVICE_SID",
			"EMAIL_HOST", "EMAIL_PORT", "EMAIL_USER", "EMAIL_PASSWORD", "IMAGEKIT_PUBLIC_KEY", "IMAGEKIT_PRIVATE_KEY", "IMAGEKIT_URL_ENDPOINT", "UPSTASH_REDIS_REST_URL", "UPSTASH_REDIS_REST_TOKEN",
		}

		// Load environment variables dynamically
		for _, key := range envVars {
			value := os.Getenv(key)
			if value == "" {
				log.Warnf("%s is not set", key)
			}

			switch key {
			case "ACCESS_TOKEN_SECRET":
				c.ACCESS_TOKEN_SECRET = value
			case "REFRESH_TOKEN_SECRET":
				c.REFRESH_TOKEN_SECRET = value
			case "PORT":
				c.PORT = value
			case "NEO4J_USER":
				c.NEO4J_USER = value
			case "NEO4J_PASSWORD":
				c.NEO4J_PASSWORD = value
			case "NEO4J_URI":
				c.NEO4J_URI = value
			case "TWILLIO_ACCOUNT_SID":
				c.TWILLIO_ACCOUNT_SID = value
			case "TWILLIO_ACCOUNT_PASSWORD":
				c.TWILLIO_ACCOUNT_PASSWORD = value
			case "TWILLIO_ACCOUNT_SERVICE_SID":
				c.TWILLIO_ACCOUNT_SERVICE_SID = value
			case "EMAIL_HOST":
				c.EMAIL_HOST = value
			case "EMAIL_PORT":
				c.EMAIL_PORT = value
			case "EMAIL_USER":
				c.EMAIL_USER = value
			case "EMAIL_PASSWORD":
				c.EMAIL_PASSWORD = value
			case "IMAGEKIT_PUBLIC_KEY":
				c.IMAGEKIT_PUBLIC_KEY = value
			case "IMAGEKIT_PRIVATE_KEY":
				c.IMAGEKIT_PRIVATE_KEY = value
			case "IMAGEKIT_URL_ENDPOINT":
				c.IMAGEKIT_URL_ENDPOINT = value
			case "UPSTASH_REDIS_REST_URL":
				c.UPSTASH_REDIS_REST_URL = value
			case "UPSTASH_REDIS_REST_TOKEN":
				c.UPSTASH_REDIS_REST_TOKEN = value
			}
		}
		c.ENVIRONMENT = environment
	}

	return c, nil
}
