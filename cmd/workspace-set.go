package cmd

import (
	"github.com/nicolasdscp/giwow/internal/workspace"
	"github.com/nicolasdscp/giwow/logger"
	"github.com/spf13/cobra"
)

var (
	workspaceSetType = ""
	workspaceSetCmd  = &cobra.Command{
		Use:   "set",
		Short: "Manage workspaces properties",
		Args:  cobra.ExactArgs(0),
		Long:  ``,
		Run:   runWorkspaceSet,
	}
)

func init() {
	workspaceCmd.AddCommand(workspaceSetCmd)

	workspaceSetCmd.PersistentFlags().StringVar(&workspaceSetType, "type", "", "Set the type of workspace, it can be default or gitlab")
}

func runWorkspaceSet(_ *cobra.Command, _ []string) {
	changeMade := false

	if workspaceSetType != "" {
		if !workspace.ValidateType(workspaceSetType) {
			logger.Print("Invalid value for type, no change made")
		}
		workspace.Current.Type = workspaceSetType
		changeMade = true
		logger.Print("Workspace type set to " + workspaceSetType)
	}

	if !changeMade {
		logger.Print("No change were made")
	}

	cobra.CheckErr(workspace.Save())
}
