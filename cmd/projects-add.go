package cmd

import (
	"github.com/nicolasdscp/giwow/internal/workspace"
	"github.com/nicolasdscp/giwow/logger"
	"github.com/spf13/cobra"
)

var (
	projectsAddClone bool
	projectsAddCmd   = &cobra.Command{
		Use:   "add [project1] [project2] ...",
		Args:  cobra.MinimumNArgs(1),
		Short: "Manage projects in the current workspace",
		Long: `Add a new project to the current workspace. 
The [name] must correspond to the git remote url.`,
		Run: runProjectAdd,
	}
)

func init() {
	projectsCmd.AddCommand(projectsAddCmd)

	projectsAddCmd.Flags().BoolVarP(&projectsAddClone, "clone", "c", false, "Directly clone the project")
}

func runProjectAdd(_ *cobra.Command, args []string) {
	for _, project := range args {
		cobra.CheckErr(workspace.AddProject(project))
		logger.Print("✅ Project %s added to the current workspace", project)
	}

	cobra.CheckErr(workspace.Save())

	if projectsAddClone {
		runProjectClone(nil, args)
	}
}
