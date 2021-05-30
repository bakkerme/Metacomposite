package main

import (
	"encoding/json"
	"fmt"
	utils "gitlab.com/hyperfocus.systems/hyperfocus-utils"
)

// Config represents application-level configuration
type Config struct {
	Feeds  []Feed
	Groups []Group
}

// FileConfigProvider provides configuration loading and parsing
type FileConfigProvider struct{}

// LoadConfig will load a config file off disk
func (cp FileConfigProvider) LoadConfig(path string) (*Config, error) {
	return cp.loadConfig(path, &utils.DirReader{})
}

func (cp FileConfigProvider) loadConfig(path string, dr utils.DirReaderProvider) (*Config, error) {
	file, err := dr.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Could not find Config file, please see readme and add a valid config to %s. Error %s", path, err)
	}

	var cfg Config
	if err := json.Unmarshal([]byte(file), &cfg); err != nil {
		return nil, fmt.Errorf("Can't unmarshal config file. Loading up %s, got error %s", path, err)
	}

	return &cfg, nil
}
