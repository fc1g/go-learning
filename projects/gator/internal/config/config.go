package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fc1g/gator/internal/types"
)

func Read() (*types.Config, error) {
	var config types.Config

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return &config, err
	}

	bytes, err := os.ReadFile(configFilePath)
	if err != nil {
		return &config, fmt.Errorf("error reading config file: %v", err)
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return &config, fmt.Errorf("error parsing json file: %v", err)
	}

	return &config, err
}

func Write(config *types.Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("error marshaling config: %v", err)
	}

	err = os.WriteFile(configFilePath, bytes, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}

	return nil
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %v", err)
	}

	return filepath.Join(homeDir, configFileName), nil
}
