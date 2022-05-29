package config

import (
	"chainlink-trial/util"
	"github.com/spf13/viper"
)

func LoadConfig(networkProvider string) (*Config, error) {
	config := Config{}
	viper.SetConfigName("network")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(util.ProjectRoot)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.UnmarshalKey(networkProvider, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

type Config struct {
	Url string `mapstructure:"url"`
	Key string `mapstructure:"key"`
}
