package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v2"
)

// Config struct
type Config struct {
	Inputs []*Input `yaml:"inputs"`
}

// Input struct
type Input struct {
	Name string `yaml:"name"`
	Config map[string]interface{} `yaml:"config"`
}

// Load will import a nbvmtagger configuration from a file
func Load(file string) *Config {
	var cfg Config

	fh, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.NewDecoder(fh).Decode(&cfg); err != nil {
		log.Fatal(err)
	}

	fh.Close()

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}