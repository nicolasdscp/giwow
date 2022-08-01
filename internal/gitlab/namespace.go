package gitlab

func GetNamespaceFullPath(groupSlug string, token string, host string) (string, error) {
	client, err := GetClient(token, host)
	if err != nil {
		return "", err
	}

	namespace, _, err := client.Namespaces.GetNamespace("giwow/veepee")
	if err != nil {
		return "", err
	}

	return namespace.FullPath, err
}
