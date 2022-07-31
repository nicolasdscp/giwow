package cmd

import (
	"os"
	"path"

	"github.com/manifoldco/promptui"
	"github.com/nicolasdscp/giwow/config"
	"github.com/nicolasdscp/giwow/internal/workspace"
	"github.com/nicolasdscp/giwow/logger"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Delete all configuration files and all workspaces",
	Long: `WARNING: This will delete all configuration files and all workspaces.
This action is irreversible.

Note that giwow will only delete his own configuration files including $HOME/.giwow and .giwow in all workspaces.`,
	Run: runPrune,
}

func init() {
	rootCmd.AddCommand(pruneCmd)
}

func runPrune(_ *cobra.Command, _ []string) {
	logger.Print("WARNING: This will delete all configuration files and all workspaces.")

	prompt := promptui.Prompt{
		Label:     "Are you sure you want to proceed",
		IsConfirm: true,
	}

	result, err := prompt.Run()
	cobra.CheckErr(err)

	if result != "y" {
		logger.Print("Aborted")
		return
	}

	for _, workspacePath := range config.Workspaces {
		if err = workspace.DeleteWorkspace(path.Join(workspacePath, ".giwow")); err != nil {
			logger.Verbose("Error deleting workspace %s: %s", workspacePath, err)
			continue
		}
		logger.Verbose("Workspace %s deleted", workspacePath)
	}

	logger.Print("%d workspace(s) deleted", len(config.Workspaces))

	config.SaveConfigFileOnExit = false
	cobra.CheckErr(os.RemoveAll(config.CfgDir))

	logger.Print("Configuration files deleted, operation successful")
}
