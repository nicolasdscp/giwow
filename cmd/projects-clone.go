package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/nicolasdscp/giwow/internal/workspace"
	"github.com/nicolasdscp/giwow/logger"
	"github.com/spf13/cobra"
)

var (
	projectsCloneCmd = &cobra.Command{
		Use:   "clone [project1] [project2] ...",
		Short: "Manage projects in the current workspace",
		Long: `Add a new project to the current workspace. 
The [name] must correspond to the git remote url.`,
		Run: runProjectClone,
	}
)

func init() {
	projectsCmd.AddCommand(projectsCloneCmd)
}

func runProjectClone(_ *cobra.Command, args []string) {
	var projectsToClone []string

	if len(args) == 0 {
		var projects []string

		prompt := survey.MultiSelect{
			Message: "Select projects to clone",
			Options: workspace.Current.Projects,
		}

		cobra.CheckErr(survey.AskOne(&prompt, &projects))

		projectsToClone = projects
	} else {
		projectsToClone = args
	}

	for _, project := range projectsToClone {
		_, err := workspace.CloneProject(project)
		if err != nil {
			logger.Print("❌ Project %s not cloned: %s", project, err.Error())
			continue
		}
	}
}
