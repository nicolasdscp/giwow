package cmd

import (
	"github.com/nicolasdscp/giwow/errors"
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
	Run:              runProjects,
}

func init() {
	rootCmd.AddCommand(projectsCmd)

	projectsCmd.PersistentFlags().String("netrc", "", `The path to the netrc file, default is $HOME/.netrc`)
}

func persistentPreRunProjects(cmd *cobra.Command, _ []string) {
	cobra.CheckErr(workspace.ResolveCurrent())
	cobra.CheckErr(netrc.ResolveCurrent(cmd.Flag("netrc").Value.String()))

	if workspace.Current == nil {
		logger.Fatal(errors.ErrWorkspaceNotFound().Error())
	}
}

func runProjects(cmd *cobra.Command, args []string) {
	cmd.Usage()
}
