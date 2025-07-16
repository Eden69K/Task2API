package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		URL     string `yaml:"url"`
		AppName string `yaml:"app_name"`
	} `yaml:"database"`

	TargetAPI struct {
		URL           string `yaml:"url"`
		Authorization string `yaml:"authorization"`
	} `yaml:"target_api"`

	RateLimit struct {
		RequestsPerSecond int `yaml:"requests_per_second"`
	} `yaml:"rate_limit"`

	Logging struct {
		Filename   string `yaml:"filename"`
		MaxSizeMB  int    `yaml:"max_size_mb"`
		MaxBackups int    `yaml:"max_backups"`
		MaxAgeDays int    `yaml:"max_age_days"`
	} `yaml:"logging"`
}

//------------------------------------------------------------------------------------------

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	expanded := os.ExpandEnv(string(data))

	var cfg Config
	if err := yaml.Unmarshal([]byte(expanded), &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func MustLoadConfig(path string) *Config {
	config, err := LoadConfig(path)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return config
}
