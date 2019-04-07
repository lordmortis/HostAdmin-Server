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

type Config struct {
	Database DatabaseConfig `json:"db"`
	Server ServerConfig `json:"server"`
}

func LoadConfig(filename string) (*Config, error) {
	fileString, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config = Config{}

	err = json.Unmarshal(fileString, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c ServerConfig) String() string {
	return c.BindAddress + ":" + strconv.Itoa(c.Port)
}