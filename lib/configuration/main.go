package configuration

import "github.com/spf13/viper"

type Configuration struct {
	Port  string
	IsDev bool
}

func InitConfig() Configuration {
	viper.SetEnvPrefix("tawny")
	viper.BindEnv("port")
	viper.SetDefault("port", "8080")

	viper.BindEnv("dev")
	viper.SetDefault("dev", true)
	return Configuration{
		Port:  viper.GetString("port"),
		IsDev: viper.GetBool("dev"),
	}
}
