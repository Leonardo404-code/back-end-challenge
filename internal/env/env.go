package env

import (
	"github.com/spf13/viper"
)

func GetInt(envName string) int {
	if !viper.IsSet(envName) {
		panic("env variable not found: " + envName)
	}
	return viper.GetInt(envName)
}

func GetString(envName string) string {
	if !viper.IsSet(envName) {
		panic("env variable not found: " + envName)
	}
	return viper.GetString(envName)
}

func GetBool(envName string) bool {
	if !viper.IsSet(envName) {
		panic("env variable not found: " + envName)
	}
	return viper.GetBool(envName)
}

func GetStringSlice(envName string) []string {
	if !viper.IsSet(envName) {
		panic("env variable not found: " + envName)
	}
	return viper.GetStringSlice(envName)
}
