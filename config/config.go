package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func setupEnvVariables(confName string) {
	viper.AutomaticEnv()
	viper.SetConfigName(confName)
	viper.AddConfigPath("$HOME/go/src/github.com/raghukul01/Remotify/config") // usual path to config
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".") // optionally look for config in the working directory
}

// Load the config, can be accessed globally
func Load() {
	logrus.Info("Func: Viper Load()")
	err := viper.BindEnv("env")
	if err != nil {
		logrus.WithError(err).Fatal("env load failed")
		panic(err)
	}

	configName := "data"
	logrus.WithField("configName", configName).Info()

	setupEnvVariables(configName)

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		logrus.WithError(err).Fatal("Fatal error config file")
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Info("Config file changed:", e.Name)
	})
}
