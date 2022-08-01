package cmd

import (
	"github.com/nicolasdscp/giwow/errors"
	"github.com/nicolasdscp/giwow/internal/gitlab"
	"github.com/nicolasdscp/giwow/internal/netrc"
	"github.com/nicolasdscp/giwow/internal/workspace"
	"github.com/nicolasdscp/giwow/logger"
	"github.com/spf13/cobra"
)

var (
	workspaceInfoCmd = &cobra.Command{
		Use:   "info",
		Short: "Display workspaces properties",
		Args:  cobra.ExactArgs(0),
		Long:  ``,
		Run:   runWorkspaceInfo,
	}
)

func init() {
	workspaceCmd.AddCommand(workspaceInfoCmd)
}

func runWorkspaceInfo(_ *cobra.Command, _ []string) {
	if workspace.Current == nil {
		logger.Fatal(errors.ErrWorkspaceNotFound().Error())
	}
	logger.Print("🔎 Information for workspace %s", workspace.CurrentWorkspaceDir)
	logger.Print("root: " + workspace.Current.Root)
	logger.Print("type: " + workspace.Current.Type)
	logger.Print("projects: %d", len(workspace.Current.Projects))
	logger.Verbose("host: %s", workspace.CurrentWorkspaceHost)
	logger.Verbose("slug: %s", workspace.CurrentWorkspaceSlug)

	if logger.VerboseMode {
		_, token := netrc.GetWorkspaceCredentials(workspace.Current.Root)
		namespacePath, err := gitlab.GetNamespaceFullPath(workspace.CurrentWorkspaceSlug, token, workspace.CurrentWorkspaceHost)
		if err != nil {
			logger.Fatal("💀 Error getting namespace full path: %s", err.Error())
		}
		logger.Verbose("namespace: %s", namespacePath)
	}
}
