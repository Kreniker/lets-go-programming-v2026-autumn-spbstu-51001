package main

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("error reading config file %q: %v", path, err))
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Sprintf("error unmarshaling YAML config %q: %v", path, err))
	}

	if cfg.InputFile == "" {
		panic("in config not set field input-file")
	}

	if cfg.OutputFile == "" {
		panic("in config not set field output-file")
	}

	return cfg
}

func main() {

}
