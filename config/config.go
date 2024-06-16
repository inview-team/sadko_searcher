package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

type ServerConfig struct {
	Port                   string
	Ssl                    bool
	EndpointVideoProcessor string
}

type PostgresConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DbName       string
	SSLMode      string
	PoolMaxConns int
}

type Config interface {
	ParseConfig() (AppConfig, error)
}

type ConfigDriver struct {
	v *viper.Viper
}

func LoadNewConfig() (Config, error) {
	v := viper.New()
	v.SetConfigFile(findConfigPath())

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config: file not found")
		}
		return nil, err
	}

	return &ConfigDriver{
		v: v,
	}, nil
}

func (c *ConfigDriver) ParseConfig() (AppConfig, error) {
	config := AppConfig{}

	err := c.v.Unmarshal(&config)
	if err != nil {
		return AppConfig{}, fmt.Errorf("config: unable to decode into struct: %w", err)
	}
	return config, nil
}

func findConfigPath() string {
	customPath := os.Getenv("CONFIG_PATH")

	configPaths := map[string]string{
		"local":  "config/config_sample.yml",
		"custom": customPath,
	}
	if configPaths["custom"] != "" {
		return configPaths["custom"]
	} else {
		return configPaths["local"]
	}
}
