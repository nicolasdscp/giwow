package errors

import "fmt"

// ErrWorkspaceAlreadyExists is triggered when a workspace is already initialized.
func ErrWorkspaceAlreadyExists() error {
	return fmt.Errorf("workspace already exists in this directory")
}

// ErrWorkspaceNotFound is triggered when a workspace is not found.
func ErrWorkspaceNotFound() error {
	return fmt.Errorf("workspace not found")
}
