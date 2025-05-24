package config

import (
	"fmt"
	"testing"
)

func TestLoadConfigViper(t *testing.T) {
	viper, err := LoadConfigViper("./config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(viper.GetString("name"))
}
