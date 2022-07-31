package cmd

import (
	"github.com/nicolasdscp/giwow/internal/workspace"
	"github.com/spf13/cobra"
)

// WorkspaceCmd represents the workspace command
var workspaceCmd = &cobra.Command{
	Use:               "workspace",
	Short:             "Manage workspaces",
	Long:              ``,
	PersistentPreRunE: persistentPreRunEWorkspace,
	Run:               runWorkspace,
}

func init() {
	rootCmd.AddCommand(workspaceCmd)
	workspaceCmd.DisableFlagsInUseLine = true
}

func persistentPreRunEWorkspace(cmd *cobra.Command, args []string) error {
	return workspace.ResolveCurrent()
}

func runWorkspace(cmd *cobra.Command, args []string) {
	cmd.Usage()
}
