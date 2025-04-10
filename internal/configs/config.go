package configs

import (
	"github.com/joaoasantana/e-inventory-service/pkg/config"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	App      config.AppInfo
	Database config.DatabaseInfo
	Kafka    config.KafkaInfo
	Server   config.ServerInfo
}

func LoadAllConfig() *Config {
	viperConfiguration()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var configs *Config

	if err := viper.Unmarshal(&configs); err != nil {
		panic(err)
	}

	configs.Database.User = os.Getenv("DB_USER")
	configs.Database.Pass = os.Getenv("DB_PASSWORD")

	return configs
}

func viperConfiguration() {
	viper.AutomaticEnv()

	viper.SetConfigName("config.debug")
	viper.AddConfigPath("./config")

	viper.SetConfigType("yaml")
}
