package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var v *viper.Viper

func LoadConfigViper(config string) (*viper.Viper, error) {
	if len(config) == 0 {
		return nil, fmt.Errorf("no config file provided")
	}
	viper.SetConfigFile(config)
	viper.AutomaticEnv()
	viper.WatchConfig()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	v = viper.GetViper()
	return v, nil
}
