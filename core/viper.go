package core

import (
	"strings"

	"github.com/spf13/viper"
)

var Viper *viper.Viper

func InitViper(filename *string) {
	components := strings.Split(*filename, ".")

	_viper := viper.New()
	_viper.SetConfigName(components[0])
	_viper.SetConfigType(components[1])
	// 默认在当前项目的config目录下寻找配置
	_viper.AddConfigPath("./config")

	if err := _viper.ReadInConfig(); err != nil {
		panic(err)
	}

	Viper = _viper
}
