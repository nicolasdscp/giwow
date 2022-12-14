package cmd

import (
	"github.com/nicolasdscp/giwow/internal/exception"
	"github.com/nicolasdscp/giwow/internal/netrc"
	"github.com/nicolasdscp/giwow/internal/workspace"
	"github.com/nicolasdscp/giwow/logger"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:              "projects",
	Short:            "Manage projects in the current workspace",
	Long:             ``,
	PersistentPreRun: persistentPreRunProjects,
	RunE:             runProjectsE,
}

func init() {
	rootCmd.AddCommand(projectsCmd)

	projectsCmd.PersistentFlags().String("netrc", "", `The path to the netrc file, default is $HOME/.netrc`)
}

// Load workspace and netrc modules
func persistentPreRunProjects(cmd *cobra.Command, _ []string) {
	cobra.CheckErr(workspace.ResolveCurrent())
	cobra.CheckErr(netrc.ResolveCurrent(cmd.Flag("netrc").Value.String()))

	if workspace.Current == nil {
		logger.Fatal(exception.ErrWorkspaceNotFound().Error())
	}
}

func runProjectsE(cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}
