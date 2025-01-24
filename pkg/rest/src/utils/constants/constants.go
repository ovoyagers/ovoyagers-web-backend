package constants

import (
	"fmt"

	"github.com/petmeds24/backend/config"
)

const Domain = "ovoyagers-web-backend.onrender.com"

type Constants struct {
	HOST            string `mapstructure:"HOST"`
	PORT            string `mapstructure:"PORT"`
	BASE_URL        string `mapstructure:"BASE_URL"`
	SCHEMA_BASE_URL string `mapstructure:"SCHEMA_BASE_URL"`
	IS_SECURE       bool   `mapstructure:"IS_SECURE"`
	SCHEMA          string `mapstructure:"SCHEMA"`
}

func GetConstants(cfg *config.Config) *Constants {
	if cfg.ENVIRONMENT == "local" {
		return &Constants{
			HOST:            "localhost",
			PORT:            cfg.PORT,
			BASE_URL:        fmt.Sprintf("%s:%s", "localhost", cfg.PORT),
			SCHEMA_BASE_URL: fmt.Sprintf("%s://%s:%s", "http", "localhost", cfg.PORT),
			IS_SECURE:       false,
			SCHEMA:          "http",
		}
	}
	return &Constants{
		HOST:            "",
		PORT:            "80",
		BASE_URL:        Domain,
		SCHEMA_BASE_URL: fmt.Sprintf("%s://%s", "https", Domain),
		IS_SECURE:       true,
		SCHEMA:          "https",
	}
}
