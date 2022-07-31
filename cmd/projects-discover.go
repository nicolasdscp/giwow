package cmd

import (
	"github.com/nicolasdscp/giwow/internal/gitlab"
	"github.com/nicolasdscp/giwow/internal/netrc"
	"github.com/nicolasdscp/giwow/internal/workspace"
	"github.com/nicolasdscp/giwow/logger"
	"github.com/spf13/cobra"
)

var (
	projectsDiscoverCmd = &cobra.Command{
		Use:   "discover",
		Args:  cobra.ExactArgs(0),
		Short: "Try to discover projects in the current workspace",
		Long:  ``,
		Run:   runProjectDiscover,
	}
)

func init() {
	projectsCmd.AddCommand(projectsDiscoverCmd)
	projectsDiscoverCmd.DisableFlagsInUseLine = true
}

func runProjectDiscover(_ *cobra.Command, _ []string) {
	logger.Print("🔎 Discovering projects in the current workspace ...")

	if workspace.Current.Type == workspace.TypeGitlab {
		logger.Print("🦊 Using Gitlab API to discover projects ...")
		_, token := netrc.GetWorkspaceCredentials(workspace.Current.Root)

		projects, err := gitlab.DiscoverProjects(workspace.CurrentWorkspaceSlug, token, workspace.CurrentWorkspaceHost)
		if err != nil {
			logger.Fatal("💀 Error discovering projects: %s", err.Error())
		}

		logger.Print("🎉 Discovered %d projects", len(projects))

		workspace.Current.Projects = projects
		cobra.CheckErr(workspace.Save())

		logger.Print("✅  %d project(s) added to the workspace", len(projects))
	}
}
