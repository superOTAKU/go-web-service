package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

type ApplicationConfig struct {
	Server  ServerConfig
	Db      DbConfig
	Logging LoggingConfig
}

type ServerConfig struct {
	Host string `default:"0.0.0.0"`
	Port int    `default:"8080"`
}

type LoggingConfig struct {
	Filepath string `default:"log/app.log"`
}

func (c *ServerConfig) GetListenAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type DbConfig struct {
	Dsn string // DataSource Name
}

func LoadConfig(path string) (*ApplicationConfig, error) {
	var config ApplicationConfig
	if err := configor.Load(&config, path); err != nil {
		return nil, err
	}
	return &config, nil
}
