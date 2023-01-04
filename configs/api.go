package configs

import "github.com/spf13/viper"

type ApiConfig struct {
	Port int
}

func LoadApi() ApiConfig {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()
	conf := ApiConfig{}
	conf.Port = conf.GetPort()
	return conf
}

func (config *ApiConfig) GetPort() int {
	viper.SetDefault("PORT", 5001)
	return viper.GetInt("PORT")
}
