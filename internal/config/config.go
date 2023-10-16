package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

func GetConf() *Config {
	c := &Config{}
	info, err := os.ReadFile("./../conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(info, c)
	return c
}
