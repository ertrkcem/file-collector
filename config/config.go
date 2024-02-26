package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config represents the configuration for a server connection.
type Config struct {
	ServerType string   `json:"server_type"` // "ftp" or "sftp"
	Host       string   `json:"host"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Path       string   `json:"path"`
	Targets    []Target `json:"targets"` // Array of target APIs
}

// Target represents the configuration for a single API target.
type Target struct {
	APIURL   string `json:"api_url"`
	APIToken string `json:"api_token"`
}

// LoadConfigurations reads the JSON configuration file and returns an array of Config structs.
func LoadConfigurations(filePath string) ([]Config, error) {
	// Read the file content
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal the JSON data
	var configs []Config
	err = json.Unmarshal(data, &configs)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config data: %w", err)
	}

	return configs, nil
}
