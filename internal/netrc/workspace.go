package netrc

import "strings"

func GetWorkspaceCredentials(workspaceRoot string) (login string, password string) {
	lastFoundMachine := &Machine{Name: ""}

	for _, machine := range Current.GetMachines() {
		if strings.Contains(workspaceRoot, machine.Name) && len(lastFoundMachine.Name) < len(machine.Name) {
			lastFoundMachine = machine
		}
	}

	return lastFoundMachine.Get("login"), lastFoundMachine.Get("password")
}
