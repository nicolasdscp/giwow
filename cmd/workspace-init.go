package cmd

import (
	"github.com/nicolasdscp/giwow/internal/workspace"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var workspaceInitCmd = &cobra.Command{
	Use:   "init [root]",
	Short: "Initialize a new workspace",
	Args:  cobra.ExactArgs(1),
	Long: `This command allows you to initialize a new workspace in the current directory.
[root] defines a git remote url of your group. It can be a github url or a gitlab url.
eg: github.com/my-group

You can also create a workspaces in child directories of the current workspace directory.
When you initialize a workspace, it will create a .giwow directory in the current path.

Note that when you run a giwow command, it will automatically work with the closest workspace`,
	Run: runWorkspaceInit,
}

func init() {
	workspaceCmd.AddCommand(workspaceInitCmd)
	workspaceInitCmd.DisableFlagsInUseLine = true
}

func runWorkspaceInit(_ *cobra.Command, args []string) {
	cobra.CheckErr(workspace.Init(args[0]))
}
