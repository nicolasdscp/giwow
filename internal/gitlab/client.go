package gitlab

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

// GetClient returns a new GitLab client.
func GetClient(token string, host string) (*gitlab.Client, error) {
	return gitlab.NewClient(token, gitlab.WithBaseURL(fmt.Sprintf("https://%s/api/v4", host)))
}
