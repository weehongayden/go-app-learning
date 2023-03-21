package config

import "github.com/spf13/viper"

type Database struct {
	Host string `mapstructure:"host"`
	Port uint16 `mapstructure:"port"`
	Name string `mapstructure:"name"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
}

type Config struct {
	Database Database `mapstructure:"database"`
}

func LoadConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
