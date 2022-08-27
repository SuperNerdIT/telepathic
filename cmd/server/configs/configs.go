package configs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Configs struct {
	Port string
	Host string
}

func init() {
	loadConfiguration()
}

func loadConfiguration() {
	configName := "default"
	if _, err := os.Stat("config.yml"); err == nil {
		configName = "config"
	} else if !errors.Is(err, os.ErrNotExist) {
		fmt.Print(err)
		panic("Unexpected error trying read config file")
	}
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")

	path := getCallerPath()
	viper.AddConfigPath(path)
}

func getCallerPath() string {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error getting the configuration path")
	}
	return filepath.Dir(b)
}

func MakeConfigs() *Configs {
	// we could trigger something to read envs or command line arguments here
	// so this could be dynamic

	return &Configs{
		Port: readFromFile("server.addr.port"),
		Host: readFromFile("server.addr.host"),
	}
}

func readFromFile(key string) string {
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		panic("Error reading config file")
	}

	return viper.GetString(key)
}
