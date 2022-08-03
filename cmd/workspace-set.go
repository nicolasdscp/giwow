package cmd

import (
	"github.com/nicolasdscp/giwow/internal/exception"
	"github.com/nicolasdscp/giwow/internal/workspace"
	"github.com/nicolasdscp/giwow/logger"
	"github.com/spf13/cobra"
)

const (
	flagWorkspaceSetType        = "type"
	flagWorkspaceSetDefaultOpen = "defaultOpen"
)

var (
	workspaceSetType        string
	workspaceSetDefaultOpen string
	workspaceSetCmd         = &cobra.Command{
		Use:   "set",
		Short: "Manage workspaces properties",
		Args:  cobra.ExactArgs(0),
		Long:  ``,
		Run:   runWorkspaceSet,
	}
)

func init() {
	workspaceCmd.AddCommand(workspaceSetCmd)

	workspaceSetCmd.Flags().StringVar(&workspaceSetType, flagWorkspaceSetType, "", "Set the type of workspace, it can be default or gitlab")
	workspaceSetCmd.Flags().StringVar(&workspaceSetDefaultOpen, flagWorkspaceSetDefaultOpen, "", "Set your default IDE to use when opening a project")
}

func runWorkspaceSet(_ *cobra.Command, _ []string) {
	if workspace.Current == nil {
		logger.Fatal(exception.ErrWorkspaceNotFound().Error())
	}

	changeMade := false

	if workspaceSetType != "" {
		if !workspace.ValidateType(workspaceSetType) {
			logger.Fatal("Invalid value for %s, no change made", flagWorkspaceSetType)
		}
		workspace.Current.Type = workspaceSetType
		changeMade = true
		logger.Print("Workspace %s set to %s", flagWorkspaceSetType, workspaceSetType)
	}

	if workspaceSetDefaultOpen != "" {
		if !workspace.ValidateDefaultOpen(workspaceSetDefaultOpen) {
			logger.Fatal("Invalid value for %s, no change made", flagWorkspaceSetDefaultOpen)
		}
		workspace.Current.DefaultOpen = workspaceSetDefaultOpen
		changeMade = true
		logger.Print("Workspace %s set to %s", flagWorkspaceSetDefaultOpen, workspaceSetDefaultOpen)
	}

	if !changeMade {
		logger.Print("No flag provided")
	}

	cobra.CheckErr(workspace.Save())
}
