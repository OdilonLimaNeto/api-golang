package config

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	config                     *configuration
	STRING_CONNECTION_DATABASE = ""
)

type configuration struct {
	API APIConfiguration
	DB  DBConfiguration
}

type APIConfiguration struct {
	PORT string
}

type DBConfiguration struct {
	USER     string
	PASSWORD string
	HOST     string
	PORT     string
	NAME     string
}

func Load() error {
	viper.SetConfigName("configurations")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("No config file found %v", err)
			return err
		}
	}

	config = new(configuration)
	config.API = APIConfiguration{
		PORT: viper.GetString("api.port"),
	}

	config.DB = DBConfiguration{
		USER:     viper.GetString("database.user"),
		PASSWORD: viper.GetString("database.password"),
		HOST:     viper.GetString("database.host"),
		PORT:     viper.GetString("database.port"),
		NAME:     viper.GetString("database.name"),
	}
	return nil
}

func GetDB() DBConfiguration {
	return config.DB
}

func GetServerPORT() string {
	return config.API.PORT
}
