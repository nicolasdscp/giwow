package workspace

import (
	"fmt"
	"os"
	"path"

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

// getProjectUrl returns the url of the current project.
func getProjectUrl(name string) string {
	return fmt.Sprintf("https://%s/%s", Current.Root, name)
}
