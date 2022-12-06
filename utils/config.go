package utils

import (
	"github.com/spf13/viper"
)

type ConfigStruct struct {
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabaseUser     string `mapstructure:"DATABASE_USER"`
	DatabasePort     string `mapstructure:"DATABASE_PORT"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	SecretKey        string `mapstructure:"SECRET_KEY"`
	JwtLife          string `mapstructure:"JWT_LIFE"`
}

var Config ConfigStruct

func Init(path string) (config ConfigStruct, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
func LoadConfig() {
	var err error
	Config, err = Init(".")
	if err != nil {
		return
	}
}
