package gitlab

func GetNamespaceFullPath(token string, host string) (string, error) {
	client, err := GetClient(token, host)
	if err != nil {
		return "", err
	}

	namespace, _, err := client.Namespaces.GetNamespace("giwow/veepee")

	return namespace.FullPath, err
}
