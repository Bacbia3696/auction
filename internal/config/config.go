package config

import "github.com/spf13/viper"

type Config struct {
	DbSource     string `mapstructure:"db_source"`
	Server       Server `mapstructure:"server"`
	Environment  string `mapstructure:"environment"`
	TokenSignKey string `mapstructure:"token_sign_key"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port uint32 `mapstructure:"port"`
}

var cfg *Config

func Load(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func Get() *Config {
	return cfg
}
