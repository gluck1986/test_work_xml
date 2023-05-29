package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// Config is a config
type Config struct {
	LogLevel         string `envconfig:"LOG_LEVEL"`
	PgURL            string `envconfig:"PG_URL"`
	PgMigrationsPath string `envconfig:"PG_MIGRATIONS_PATH"`
	HTTPAddr         string `envconfig:"HTTP_ADDR"`
	SdnXMLSource     string `envconfig:"SDN_XML_SOURCE"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment. Once.
func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal("fatal", "config", err)
		}
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal("fatal", "config", err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})
	return &config
}
