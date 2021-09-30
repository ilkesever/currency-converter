package config

import (
	"currency-converter/common/services"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

//used to store the port value and environment.
type WebServerConfig struct {
	Port        string `required:"true" split_words:"true" default:"4000"`
	Environment string `required:"true" split_words:"true" default:"dev"`
	Service     *services.ServiceConfig
}

// FromEnv pulls config from the environment file
func FromEnv() (cfg *WebServerConfig, err error) {
	cfgFilename := "../../config/config.dev.env"

	err = godotenv.Load(cfgFilename)
	if err != nil {
		fmt.Printf("No config files found. With error: %s", err.Error())
		return nil, err
	}

	cfg = &WebServerConfig{}

	err = envconfig.Process("", cfg)
	if err != nil {

		return nil, err
	}

	return cfg, nil
}
