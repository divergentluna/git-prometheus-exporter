package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadGitHubConfig(path string) (*GitHubExporterConfig, error) {
	var config GitHubExporterConfig
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading YAML file: %s", err)
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %s", err)
		return nil, err
	}

	// Overwrite GitHubToken from environment variable if set
	if envToken, exists := os.LookupEnv("GITHUB_TOKEN"); exists {
		config.GitHubToken = envToken
	}

	return &config, nil
}
