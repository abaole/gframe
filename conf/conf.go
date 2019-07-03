package conf

import "github.com/spf13/viper"

var App *AppConfig

type AppConfig struct {
	AppID string
	Viper *viper.Viper
}

func SetAppID(id string) {
	if App != nil {
		App.AppID = id
		return
	}
	App = &AppConfig{AppID: id}
}

func SetViper(viper *viper.Viper) {
	if App != nil {
		App.Viper = viper
		return
	}
	App = &AppConfig{Viper: viper}
}
