package config

import "github.com/spf13/viper"

type Config struct {
	ClickhouseURI      string `mapstructure:"CLICKHOUSE_URI"`
	ClickhousePassword string `mapstructure:"CLICKHOUSE_PASSWORD"`
	ClickhouseDatabase string `mapstructure:"CLICKHOUSE_DATABASE"`
	CloudinaryCloud    string `mapstructure:"CLOUDINARY_CLOUD"`
	CloudinaryPrivate  string `mapstructure:"CLOUDINARY_PRIVATE"`
	CloudinaryPublic   string `mapstructure:"CLOUDINARY_PUBLIC"`
	AccessTokenSecret  string `mapstructure:"ACCESS_SECRET"`
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
