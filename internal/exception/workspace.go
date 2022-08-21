package exception

import "fmt"

// ErrWorkspaceAlreadyExists is triggered when a workspace is already initialized.
func ErrWorkspaceAlreadyExists() error {
	return fmt.Errorf("workspace already exists in this directory")
}

// ErrWorkspaceNotFound is triggered when a workspace is not found.
func ErrWorkspaceNotFound() error {
	return fmt.Errorf("workspace not found")
}

// ErrProjectAlreadyExists is triggered when a project is already initialized.
func ErrProjectAlreadyExists() error {
	return fmt.Errorf("project already exists in this workspace")
}

// ErrNetrcFileNotExist is triggered when the netrc file does not exist.
func ErrNetrcFileNotExist(err error) error {
	return fmt.Errorf("netrc file does not exist: %w", err)
}
