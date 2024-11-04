package appConfig

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	configUtils "github.com/grandminingpool/pool-api/internal/common/utils/config"
	"github.com/spf13/viper"
)

type PoolAPIConfig struct {
	RequestTimeout int `mapstructure:"requestTimeout"`
}

type Config struct {
	Host    string        `mapstructure:"host"`
	Port    int16         `mapstructure:"port"`
	PoolAPI PoolAPIConfig `mapstructure:"poolAPI"`
}

const configName = "app"

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func New(configsPath string, validate *validator.Validate) (*Config, error) {
	appViper := viper.New()
	appViper.AddConfigPath(fmt.Sprintf("%s/app", configsPath))
	appViper.SetConfigType("yaml")

	appViper.SetDefault("host", "127.0.0.1")
	appViper.SetDefault("port", 5432)
	appViper.SetDefault("poolAPI.requestTimeout", 2)

	if err := configUtils.ReadConfig(appViper, configName); err != nil {
		return nil, err
	}

	config, err := configUtils.LoadConfig[Config](appViper, validate, configName)
	if err != nil {
		return nil, err
	}

	return config, nil
}
