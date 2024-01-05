package env

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func init() {
	cwd, _ := os.Getwd()
	var pathTransversal string

	for {
		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {
			resultPath := filepath.Join(cwd) + "/config/env"
			viper.AddConfigPath(resultPath)
			break
		}
		cwd = filepath.Dir(cwd)
		pathTransversal = filepath.Join(pathTransversal, "..")
		if cwd == "/" || cwd == "\\" {
			log.Printf("Couldn`t find go.mod file")
			break
		}
	}

	viper.AddConfigPath("./config/env")
	viper.SetConfigName("envs.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Couldn't read config file. Error: %v\n", err.Error())
	}
	viper.AutomaticEnv()
}
