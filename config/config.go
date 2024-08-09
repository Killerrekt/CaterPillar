package config

import "github.com/spf13/viper"

type Config struct {
	ScyllaHost        string `mapstructure:"SCYLLA_HOST"`
	ScyllaPort        string `mapstructure:"SCYLLA_PORT"`
	ScyllaKeyspace    string `mapstructure:"SCYLLA_KEYSPACE"`
	CloudinaryCloud   string `mapstructure:"CLOUDINARY_CLOUD"`
	CloudinaryPrivate string `mapstructure:"CLOUDINARY_PRIVATE"`
	CloudinaryPublic  string `mapstructure:"CLOUDINARY_PUBLIC"`
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
