package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func New() *Config {
	return &Config{}
}

type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

func (c *Config) GetConf() *Config {
	info, err := os.ReadFile("./conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(info, c)
	if err != nil {
		fmt.Fprint(os.Stderr, "error in GetConfig when unmarshaling data")
		os.Exit(1)
	}
	return c
}
