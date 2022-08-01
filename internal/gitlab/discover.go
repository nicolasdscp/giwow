package gitlab

import (
	"sort"
	"strings"

	"github.com/xanzy/go-gitlab"
)

func DiscoverProjects(groupSlug string, token string, host string) ([]string, error) {
	slugSplitLen := len(strings.Split(groupSlug, "/"))
	client, err := GetClient(token, host)
	if err != nil {
		return nil, err
	}

	currentPage := 1
	projects := make([]*gitlab.Project, 0)

	// We iterate over all pages to get all projects
	for currentPage > 0 {
		resProjects, res, reqErr := client.Groups.ListGroupProjects(groupSlug, &gitlab.ListGroupProjectsOptions{
			IncludeSubGroups: gitlab.Bool(true),
			Archived:         gitlab.Bool(false), // TODO: support archived projects with a flag (--archived)
			ListOptions: gitlab.ListOptions{
				Page: currentPage,
			},
		})
		if reqErr != nil {
			return nil, reqErr
		}

		projects = append(projects, resProjects...)
		// res.NextPage is 0 if there is no next page
		currentPage = res.NextPage
	}

	var projectNames []string
	for _, project := range projects {
		projectPathSpt := strings.Split(project.PathWithNamespace, "/")
		toAppend := strings.Join(projectPathSpt[slugSplitLen:], "/")
		// We want to avoid "" projects
		if toAppend == "" {
			continue
		}
		projectNames = append(projectNames, toAppend)
	}

	sort.Strings(projectNames)

	return projectNames, nil
}
