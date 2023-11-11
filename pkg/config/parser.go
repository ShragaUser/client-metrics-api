package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type InputMetric struct {
	Name        string              `yaml:"name"`
	Type        string              `yaml:"type"`
	Labels      []string            `yaml:"labels"`
	Description string              `yaml:"description"`
	Buckets     []float64           `yaml:"buckets,omitempty"`
	Objectives  map[float64]float64 `yaml:"objectives,omitempty"`
}

type InputFile struct {
	Metrics []InputMetric `yaml:"metrics"`
}

func ParseConfigFileFromFilePath(path string) (*InputFile, error) {
	input := &InputFile{}
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(f, input); err != nil {
		return nil, err
	}

	if input.Metrics == nil || len(input.Metrics) == 0 {
		return nil, fmt.Errorf("no metrics found in the configuration file")
	}

	return input, nil
}
