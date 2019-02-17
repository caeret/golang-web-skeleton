package app

import (
	"os"

	"github.com/spf13/viper"
)

var Config config

type config struct {
	Debug bool
	DSN   string
}

func LoadConfig() error {
	v := viper.New()
	v.SetEnvPrefix("app")
	v.AutomaticEnv()
	initConfigDefaults(v)
	{
		_, err := os.Stat("./app.yaml")
		if err == nil {
			v.SetConfigFile("./app.yaml")
			if err := v.ReadInConfig(); err != nil {
				return err
			}
		}
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	return nil
}

func initConfigDefaults(v *viper.Viper) {
	v.SetDefault("debug", false)
	v.SetDefault("dsn", "root:root@tcp(mysql:3306)/skeleton?charset=utf8mb4")
}
