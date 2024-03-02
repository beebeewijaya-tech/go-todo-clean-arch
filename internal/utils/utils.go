package utils

import "github.com/spf13/viper"

type Utils struct {
	config *viper.Viper
}

func NewUtils(config *viper.Viper) *Utils {
	return &Utils{
		config: config,
	}
}
