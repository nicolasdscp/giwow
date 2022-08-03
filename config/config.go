package config

import (
	"fmt"
	"os"
	"path"

	"github.com/nicolasdscp/giwow/internal/exception"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configType     = "json"
	configFileName = "config"
)

var (
	CfgFilePath          string
	CfgDir               string
	CurrentWd            string
	HomeDir              string
	SaveConfigFileOnExit = true
)

func Init() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	HomeDir = home

	CurrentWd, err = os.Getwd()
	cobra.CheckErr(err)

	CfgDir = path.Join(HomeDir, ".giwow")
	CfgFilePath = path.Join(CfgDir, configFileName+"."+configType)

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
		fmt.Println(exception.NoConfigFile())
	}

	loadValues()
}

func Write() error {
	if SaveConfigFileOnExit {
		setValues()
		return viper.WriteConfig()
	}

	return nil
}

func CreateConfigFile() error {
	return os.WriteFile(CfgFilePath, []byte("{}"), 0644)
}

func FileExists() bool {
	_, err := os.Stat(CfgFilePath)
	return err == nil
}

func initBaseConfig() {
	viper.SetDefault(KeyWorkspaces, []string{})
}
