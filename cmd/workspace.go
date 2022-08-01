package cmd

import (
	"github.com/nicolasdscp/giwow/internal/netrc"
	"github.com/nicolasdscp/giwow/internal/workspace"
	"github.com/spf13/cobra"
)

// WorkspaceCmd represents the workspace command
var workspaceCmd = &cobra.Command{
	Use:              "workspace",
	Short:            "Manage workspaces",
	Long:             ``,
	PersistentPreRun: persistentPreRunEWorkspace,
	Run:              runWorkspace,
}

func init() {
	rootCmd.AddCommand(workspaceCmd)
	workspaceCmd.DisableFlagsInUseLine = true

	workspaceCmd.PersistentFlags().String("netrc", "", `The path to the netrc file, default is $HOME/.netrc`)
}

func persistentPreRunEWorkspace(cmd *cobra.Command, args []string) {
	cobra.CheckErr(workspace.ResolveCurrent())
	cobra.CheckErr(netrc.ResolveCurrent(cmd.Flag("netrc").Value.String()))
}

func runWorkspace(cmd *cobra.Command, args []string) {
	cmd.Usage()
}
