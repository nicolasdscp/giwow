package config

import (
	"fmt"
	"os"
	"path"

	"github.com/nicolasdscp/giwow/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configType     = "json"
	configFileName = "config"
)

var (
	CfgPath   string
	CurrentWd string
	HomeDir   string
)

func Init() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	HomeDir = home

	CurrentWd, err = os.Getwd()
	cobra.CheckErr(err)

	CfgPath = path.Join(HomeDir, ".giwow", configFileName+"."+configType)

	viper.AddConfigPath(path.Join(HomeDir, ".giwow"))
	viper.SetConfigType(configType)
	viper.SetConfigName(configFileName)

	if !FileExists() {
		initBaseConfig()
		cobra.CheckErr(os.MkdirAll(path.Join(HomeDir, ".giwow"), 0755))
		cobra.CheckErr(CreateConfigFile())
		cobra.CheckErr(viper.WriteConfig())
	}

	if err = viper.ReadInConfig(); err != nil {
		fmt.Println(errors.NoConfigFile())
	}

	loadValues()
}

func Write() error {
	setValues()
	return viper.WriteConfig()
}

func CreateConfigFile() error {
	return os.WriteFile(CfgPath, []byte("{}"), 0644)
}

func FileExists() bool {
	_, err := os.Stat(CfgPath)
	return err == nil
}

func initBaseConfig() {
	viper.SetDefault(KeyWorkspaces, []string{})
}
