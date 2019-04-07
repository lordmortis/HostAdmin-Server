package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type DatabaseConfig struct {
	Hostname string
	Username string
	Password string
}

type ServerConfig struct {
	BindAddress string
	Port int
}

type LoggingConfig struct {
	Level string
}

type Config struct {
	Database DatabaseConfig `json:"db"`
	Server ServerConfig
	Logging LoggingConfig
}

func defaultConfig() Config {
	var config = Config{}
	config.Database.Hostname = "127.0.0.1"
	config.Server.BindAddress = "127.0.0.1"
	config.Server.Port = 3000
	config.Logging.Level = "info"
	return config
}

func LoadConfig(filename string) (*Config, error) {
	fileString, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config = defaultConfig()

	err = json.Unmarshal(fileString, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c ServerConfig) String() string {
	return c.BindAddress + ":" + strconv.Itoa(c.Port)
}