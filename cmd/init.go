package cmd

import (
	"github.com/nicolasdscp/giwow/insternal/workspace"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var workspaceInitCmd = &cobra.Command{
	Use:   "init [name]",
	Short: "Initialize a new workspace",
	Args:  cobra.ExactArgs(1),
	Long:  ``,
	Run:   runWorkspaceInit,
}

func init() {
	workspaceCmd.AddCommand(workspaceInitCmd)
	workspaceCmd.DisableFlagsInUseLine = true
}

func runWorkspaceInit(_ *cobra.Command, args []string) {
	cobra.CheckErr(workspace.Init(args[0]))
}
