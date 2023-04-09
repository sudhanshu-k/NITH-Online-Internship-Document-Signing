package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	DATABASE_URL string `mapstructure:"DATABASE_URL"`

	PORT string `mapstructure:"PORT"`

	RedisURL string `mapstructure:"REDIS_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
