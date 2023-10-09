package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	APIPort        string `mapstructure:"APIPort"`
}

// Парсим конфиг файл с помощью вайпера
func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		logrus.Error("Couldn't parse .env file")
	}

	err = viper.Unmarshal(&config)
	logrus.Info("Successfully parse .env file")
	return
}
