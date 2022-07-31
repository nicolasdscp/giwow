package workspace

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/nicolasdscp/giwow/internal/netrc"
	"github.com/nicolasdscp/giwow/logger"
)

// CloneProject clones a project from a remote url.
func CloneProject(name string) (repo *git.Repository, err error) {
	logger.Print("⬇️ Cloning into %s ...", name)

	user, pass := getCredentials()
	return git.PlainClone(path.Join(CurrentWorkspaceDir, name), false, &git.CloneOptions{
		URL:      getProjectUrl(name),
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: user,
			Password: pass,
		},
	})
}

func getCredentials() (string, string) {
	lastFoundMachine := &netrc.Machine{Name: ""}

	for _, machine := range netrc.Current.GetMachines() {
		if strings.Contains(Current.Root, machine.Name) && len(lastFoundMachine.Name) < len(machine.Name) {
			lastFoundMachine = machine
		}
	}

	return lastFoundMachine.Get("login"), lastFoundMachine.Get("password")
}

// getProjectUrl returns the url of the current project.
func getProjectUrl(name string) string {
	return fmt.Sprintf("https://%s/%s", Current.Root, name)
}
