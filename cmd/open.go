package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/nicolasdscp/giwow/internal/exception"
	"github.com/nicolasdscp/giwow/internal/workspace"
	"github.com/nicolasdscp/giwow/logger"
	"github.com/spf13/cobra"
)

var (
	openForce string
	openCmd   = &cobra.Command{
		Use:              "open [project1] [project2] ...",
		Args:             cobra.MinimumNArgs(1),
		Short:            "Open a project with the defined editor",
		Long:             ``,
		PersistentPreRun: persistentPreRunOpen,
		Run:              runOpen,
	}
)

func init() {
	rootCmd.AddCommand(openCmd)

	openCmd.Flags().StringVarP(&openForce, "force", "f", "", "Force the editor to be used")
}

// Load workspace and netrc modules
func persistentPreRunOpen(cmd *cobra.Command, _ []string) {
	cobra.CheckErr(workspace.ResolveCurrent())

	if workspace.Current == nil {
		logger.Fatal(exception.ErrWorkspaceNotFound().Error())
	}
}

func runOpen(_ *cobra.Command, args []string) {
	if workspace.Current.DefaultOpen == "" && openForce == "" {
		logger.Print("❌ No default editor set")
		return
	}

	if openForce != "" && !workspace.ValidateDefaultOpen(openForce) {
		logger.Print("❌ Invalid open method %s", openForce)
		return
	}

	if openForce == "" {
		openForce = workspace.Current.DefaultOpen
	}

	// First we prepare the list of projects to open
	projectsToOpen := make([]string, 0)
	for _, projectAsked := range args {
		associatedProject := workspace.AssociateProjectShortcut(projectAsked)
		// If we are not able to find a project with the given shortcut, we continue
		if len(associatedProject) == 0 {
			logger.Print("❌ Project %s not found", projectAsked)
			continue
		}

		// If we find more than one project with the same shortcut, we ask the user to choose
		if len(associatedProject) > 1 {
			prompt := survey.MultiSelect{
				Message: fmt.Sprintf("Many projects found for '%s'. Select project(s) you want to open:", projectAsked),
				Options: associatedProject,
			}

			var userChoice []string
			cobra.CheckErr(survey.AskOne(&prompt, &userChoice))
			projectsToOpen = append(projectsToOpen, userChoice...)
		} else {
			projectsToOpen = append(projectsToOpen, associatedProject[0])
		}
	}

	// Then we open the projects
	for _, project := range projectsToOpen {
		if err := workspace.OpenProject(project, openForce); err != nil {
			logger.Print("❌ Project %s not opened: %s", project, err.Error())
			continue
		}
	}

	cobra.CheckErr(workspace.Save())

	if projectsAddClone {
		runProjectClone(nil, args)
	}
}
