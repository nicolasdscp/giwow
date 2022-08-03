package workspace

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/nicolasdscp/giwow/internal/netrc"
	"github.com/nicolasdscp/giwow/logger"
)

// CloneProject clones a project from a remote url.
func CloneProject(name string) (*git.Repository, error) {
	logger.Print("⬇️ Cloning into %s ...", name)

	err := os.MkdirAll(path.Join(CurrentWorkspaceDir, name), 0755)
	if err != nil {
		return nil, err
	}

	user, pass := netrc.GetWorkspaceCredentials(Current.Root)
	return git.PlainClone(path.Join(CurrentWorkspaceDir, name), false, &git.CloneOptions{
		URL:      getProjectUrl(name),
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: user,
			Password: pass,
		},
	})
}

// OpenProject opens a project in the default editor.
func OpenProject(name string, openMethod string) error {
	if openMethod == "" {
		openMethod = Current.DefaultOpen
	}

	logger.Debug("Command %s %s", DefaultOpenOptions[openMethod], getProjectFullPath(name))

	cmd := exec.Command(DefaultOpenOptions[openMethod], getProjectFullPath(name))

	return cmd.Run()
}

// IsProjectExists checks if a project exists in the current workspace.
func IsProjectExists(name string) bool {
	for _, project := range Current.Projects {
		if project == name {
			return true
		}
	}

	return false
}

// AssociateProjectShortcut tries associates a shortcut to projects.
// The idea behind is to allow a user to use shortcuts to open a project.
// For example, if a user has a project named "sub/my-project",
// We want to be able to open it with "s/m" shortcut. Obviously, if
// another project is named "sub/my-project2", it will return the list of the two projects
// corresponding to the shortcut.
func AssociateProjectShortcut(shortcut string) []string {
	projects := make([]string, 0)
	shortcutSplit := strings.Split(shortcut, "/")
	// First iterate over all projects
	for _, project := range Current.Projects {
		// if the user give "my-project" as shortcut and a project matches, we add it to the list
		if len(shortcutSplit) == 1 && project == shortcutSplit[0] {
			return []string{project}
		}

		// Otherwise, we check with more "complex" shortcuts ("sub/my-project")
		projectSplit := strings.Split(project, "/")
		// We only check if the number of parts is the same
		if len(shortcutSplit) == len(projectSplit) {
			for i, s := range shortcutSplit {
				// The shortcut must match at least the prefix of the project
				if !strings.HasPrefix(projectSplit[i], s) {
					break
				}
				// If no break has been found, the shortcut is a match
				if i == len(shortcutSplit)-1 {
					projects = append(projects, project)
				}
			}
		}
	}

	return projects
}

// getProjectUrl returns the url of the current project.
func getProjectUrl(name string) string {
	return fmt.Sprintf("https://%s/%s", Current.Root, name)
}

func getProjectFullPath(name string) string {
	return path.Join(CurrentWorkspaceDir, name)
}
