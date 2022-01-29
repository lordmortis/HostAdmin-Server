package config

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type AuthConfig struct {
	SessionExpiry int
}

type DatabaseConfig struct {
	Hostname string
	Port     int
	Database string
	Username string
	Password string
}

type RedisConfig struct {
	Hostname  string
	Port      int
	Database  int
	Namespace string
	Password  string
}

type ServerConfig struct {
	BindAddress    string
	AllowedOrigins []string
	Port           int
	SessionExpiry  int
}

type LoggingConfig struct {
	Level string
}

type Config struct {
	Auth        AuthConfig
	Development bool
	Database    DatabaseConfig `json:"db"`
	Redis       RedisConfig
	Server      ServerConfig
	Logging     LoggingConfig
}

func defaultConfig() Config {
	var config = Config{}
	config.Database.Hostname = "127.0.0.1"
	config.Database.Port = 5432
	config.Redis.Hostname = "127.0.0.1"
	config.Redis.Port = 6379
	config.Redis.Database = 0
	config.Server.BindAddress = "127.0.0.1"
	config.Server.AllowedOrigins = []string{"http://localhost:3001"}
	config.Server.Port = 3000
	config.Auth.SessionExpiry = 60
	config.Logging.Level = "info"
	return config
}

func Load(filename *string) (*Config, error) {
	fileString, err := ioutil.ReadFile(*filename)
	if err != nil {
		return nil, err
	}

	var config = defaultConfig()

	err = json.Unmarshal(fileString, &config)
	if err != nil {
		return nil, err
	}

	if !config.Development {
		config.Auth.SessionExpiry = 60
	} else {
		config.Auth.SessionExpiry = 3600
	}

	return &config, nil
}

func (c ServerConfig) String() string {
	return c.BindAddress + ":" + strconv.Itoa(c.Port)
}
