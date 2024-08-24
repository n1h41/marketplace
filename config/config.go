package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Server *Server
	Db     *Db
}

type Server struct {
	Port int
}

type Db struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SslMode  string
}

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
	})

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&configInstance); err != nil {
		panic(err)
	}

	return configInstance
}
