package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Endpoints       []string `yaml:"endpoints"`
	UseClusterNodes bool     `yaml:"use_cluster_nodes"`
}

func LoadConfig() (*Config, error) {
	// Read the YAML file
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	// Parse the YAML file into the Config struct
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
