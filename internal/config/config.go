package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Config struct {
	URLs    []string `json:"urls"`
	Timeout string   `json:"timeout"`
}

func LoadConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration(cfg.Timeout)
	if err != nil {
		return nil, fmt.Errorf("invalid timeout format: %v", err)
	}
	cfg.Timeout = duration.String()
	if len(cfg.URLs) == 0 {
		return nil, errors.New("config validation failed: URLs list cannot be empty")

	}

	return &cfg, nil

}
