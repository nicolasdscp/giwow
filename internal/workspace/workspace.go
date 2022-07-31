package workspace

import (
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/nicolasdscp/giwow/config"
	"github.com/nicolasdscp/giwow/errors"
	"github.com/nicolasdscp/giwow/logger"
)

type Workspace struct {
	Name string `json:"name"`
}

var (
	Current             *Workspace
	CurrentWorkspaceDir string
)

// Init initializes a new workspace and saves it to the filesystem
func Init(name string) error {
	if CurrentWorkspaceDir == config.CurrentWd {
		return errors.ErrWorkspaceAlreadyExists()
	}
	logger.Debug("Initializing workspace")
	Current = &Workspace{Name: name}
	CurrentWorkspaceDir = config.CurrentWd

	if err := Save(); err != nil {
		return err
	}

	config.Workspaces = append(config.Workspaces, CurrentWorkspaceDir)

	return nil
}

// Save saves the workspace to the filesystem
func Save() error {
	if err := os.MkdirAll(path.Join(CurrentWorkspaceDir, ".giwow"), 0755); err != nil {
		return err
	}

	fileContent, err := json.MarshalIndent(Current, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(CurrentWorkspaceDir, ".giwow", "workspace.json"), fileContent, 0644)
}

func ResolveCurrent() error {
	lastFoundW := ""

	for _, wPath := range config.Workspaces {
		if strings.Contains(config.CurrentWd, wPath) {
			// This way, we ensure that the current workspace is closer to the current directory
			if len(wPath) > len(lastFoundW) {
				logger.Debug("Found workspace " + wPath)
				lastFoundW = wPath
			}
		}
	}

	// No workspace found, just return
	if lastFoundW == "" {
		logger.Debug("No workspace found")
		return nil
	}

	CurrentWorkspaceDir = lastFoundW

	// Otherwise, load the workspace
	return loadWorkspace(lastFoundW)
}

func DeleteWorkspace(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}

	return nil
}

func loadWorkspace(wPath string) error {
	data, err := os.ReadFile(path.Join(wPath, ".giwow", "workspace.json"))
	if err != nil {
		return err
	}

	w := &Workspace{}
	if err = json.Unmarshal(data, w); err != nil {
		return err
	}

	Current = w

	return nil
}
