package api

import (
	"encoding/json"
	"fmt"

	utils "github.com/bakkerme/hyperfocus-utils"
	"github.com/bakkerme/metacomposite/v2/types"
)

// Credentials can contain a random set of credentials for various services
type Credentials struct {
	Type   string
	Values map[string]string
}

// Config represents application-level configuration
type Config struct {
	Feeds       []types.Feed
	Groups      []types.Group
	Credentials []Credentials
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
		return nil, fmt.Errorf("could not find Config file, please see readme and add a valid config to %s. Error %s", path, err)
	}

	var cfg Config
	if err := json.Unmarshal([]byte(file), &cfg); err != nil {
		return nil, fmt.Errorf("can't unmarshal config file. Loading up %s, got error %s", path, err)
	}

	return &cfg, nil
}
