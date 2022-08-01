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

	currentPage := 1
	projects := make([]*gitlab.Project, 0)

	for currentPage > 0 {
		resProjects, res, reqErr := client.Groups.ListGroupProjects(groupSlug, &gitlab.ListGroupProjectsOptions{
			IncludeSubGroups: &boolPtr,
			OrderBy:          gitlab.String("path"),
			Sort:             gitlab.String("asc"),
			ListOptions: gitlab.ListOptions{
				Page: currentPage,
			},
		})
		if reqErr != nil {
			return nil, reqErr
		}

		projects = append(projects, resProjects...)
		currentPage = res.NextPage
	}

	var projectNames []string
	for _, project := range projects {
		projectPathSpt := strings.Split(project.PathWithNamespace, "/")
		projectNames = append(projectNames, strings.Join(projectPathSpt[slugSplitLen:], "/"))
	}

	return projectNames, nil
}
