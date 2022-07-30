/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/nicolasdscp/giwow/config"
	"github.com/nicolasdscp/giwow/logger"
	"github.com/spf13/cobra"
)

var (
	logLevel string
	rootCmd  = &cobra.Command{
		Use:   "giwow",
		Short: "Git workspace manager",
		Long:  ``,
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return config.Write()
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(config.Init)
	cobra.OnInitialize(logger.Init)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolVar(&logger.DebugMode, "debug", false, "Enable debug and verbose messages, use in development only")
	rootCmd.PersistentFlags().BoolVar(&logger.VerboseMode, "verbose", false, "Enable verbose messages")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.SuggestionsMinimumDistance = 1
}
