package conf

var App *AppConfig

type AppConfig struct {
	AppID  string
}

func SetAppID(id string)  {
	App = &AppConfig{id}
}