package config

import "github.com/spf13/viper"

type Config struct {
	ClickhouseURI      string `mapstructure:"CLICKHOUSE_URI"`
	ClickhousePassword string `mapstructure:"CLICKHOUSE_PASSWORD"`
	CloudinaryCloud    string `mapstructure:"CLOUDINARY_CLOUD"`
	CloudinaryPrivate  string `mapstructure:"CLOUDINARY_PRIVATE"`
	CloudinaryPublic   string `mapstructure:"CLOUDINARY_PUBLIC"`
}

var ConfigData Config

func LoadConfig(path string) (err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&ConfigData)
	if err != nil {
		return
	}

	return
}
