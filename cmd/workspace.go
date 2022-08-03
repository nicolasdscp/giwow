package cmd

import (
	"github.com/nicolasdscp/giwow/internal/workspace"
	"github.com/spf13/cobra"
)

// WorkspaceCmd represents the workspace command
var workspaceCmd = &cobra.Command{
	Use:              "workspace",
	Short:            "Manage workspaces",
	Long:             ``,
	PersistentPreRun: persistentPreRunEWorkspace,
	RunE:             runWorkspaceE,
}

func init() {
	rootCmd.AddCommand(workspaceCmd)
	workspaceCmd.DisableFlagsInUseLine = true
}

func persistentPreRunEWorkspace(_ *cobra.Command, _ []string) {
	cobra.CheckErr(workspace.ResolveCurrent())
}

func runWorkspaceE(cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}
