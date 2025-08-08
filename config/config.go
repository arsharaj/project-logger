package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// ElasticSearchConfig holds ES connection details
type ElasticSearchConfig struct {
	Url   string `mapstructure:"url"`
	Index string `mapstructure:"index"`
}

// Config holds the full app configuration
type Config struct {
	LogLevel      string              `mapstructure:"log_level"`
	LogFiles      []string            `mapstructure:"log_files"`
	ElasticSearch ElasticSearchConfig `mapstructure:"elasticsearch"`
}

// LoadConfig loads configuration from file and env vars
func LoadConfig(configPath string) (*Config, error) {
	v := viper.New()

	// Config file settings
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(configPath)

	// Environment variables override
	v.SetEnvPrefix("PROJECT_LOGGER")
	v.AutomaticEnv()

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal into struct
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	// Environment overrides
	if envUrl := os.Getenv("PROJECT_LOGGER_ELASTICSEARCH_URL"); envUrl != "" {
		cfg.ElasticSearch.Url = envUrl
	}
	if envIndex := os.Getenv("PROJECT_LOGGER_ELASTICSEARCH_INDEX"); envIndex != "" {
		cfg.ElasticSearch.Index = envIndex
	}

	return &cfg, nil
}
