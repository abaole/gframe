package framework

import "github.com/spf13/viper"

func InitConfig(confPath string) *viper.Viper {
	v := viper.New()
	v.SetConfigType("toml")   // 配置文件的类型
	v.AddConfigPath(confPath) // 配置文件的路径

	err := v.ReadInConfig()

	if err != nil {
		panic(err)
	}

	return v
}