package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/nicolasdscp/giwow/internal/workspace"
	"github.com/nicolasdscp/giwow/logger"
	"github.com/spf13/cobra"
)

var (
	projectsPruneCmd = &cobra.Command{
		Use:   "prune",
		Short: "Remove all projects from the current workspace",
		Long:  ``,
		Run:   runProjectPrune,
	}
)

func init() {
	projectsCmd.AddCommand(projectsPruneCmd)

	projectsPruneCmd.DisableFlagsInUseLine = true
}

func runProjectPrune(_ *cobra.Command, args []string) {
	confirm := false
	prompt := &survey.Confirm{
		Message: "Are you sure you want to remove all projects from the current workspace?",
	}
	err := survey.AskOne(prompt, &confirm)
	cobra.CheckErr(err)

	if !confirm {
		logger.Print("Projects not removed")
		return
	}

	workspace.Current.Projects = nil
	cobra.CheckErr(workspace.Save())

	logger.Print("Projects removed")
}
