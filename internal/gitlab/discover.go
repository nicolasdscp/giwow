package gitlab

import (
	"strings"

	"github.com/xanzy/go-gitlab"
)

var (
	boolPtr = true
)

func DiscoverProjects(groupSlug string, token string, host string) ([]string, error) {
	slugSplitLen := len(strings.Split(groupSlug, "/"))
	client, err := GetClient(token, host)
	if err != nil {
		return nil, err
	}

	projects, _, err := client.Groups.ListGroupProjects(groupSlug, &gitlab.ListGroupProjectsOptions{
		IncludeSubGroups: &boolPtr,
	})
	if err != nil {
		return nil, err
	}

	var projectNames []string
	for _, project := range projects {
		projectPathSpt := strings.Split(project.PathWithNamespace, "/")
		projectNames = append(projectNames, strings.Join(projectPathSpt[slugSplitLen:], "/"))
	}

	return projectNames, nil
}
