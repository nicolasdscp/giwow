package config

import "github.com/spf13/viper"

var Workspaces []string

func loadValues() {
	Workspaces = viper.GetStringSlice(KeyWorkspaces)
}

func setValues() {
	viper.Set(KeyWorkspaces, Workspaces)
}
