package config

import "github.com/spf13/viper"

var (
	ConfigName = "config"
)

type Config struct {
}

func (c Config) readConfigFile() *viper.Viper {
	viper.SetConfigName(ConfigName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config File Not Found")
		} else {
			panic("Config File Error")
		}
	}
	return viper.GetViper()
}

func (c Config) Get(key string) string {
	return c.readConfigFile().GetString(key)
}

func New() Config {
	return Config{}
}
