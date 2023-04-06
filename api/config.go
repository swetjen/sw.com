package main

import (
	"gopkg.in/yaml.v3"
)

type Config struct {
	Environment string `yaml:"env"`
	AdminEmail  string `yaml:"admin_email"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`

	DbName string `yaml:"db_name"`
	DbUri  string `yaml:"db_uri"`
}

func NewConfig(file []byte) (*Config, error) {
	var cfg = Config{}
	err := yaml.Unmarshal(file, &cfg)
	if err != nil {
		return &cfg, err
	}
	return &cfg, nil
}
