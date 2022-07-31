package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/nicolasdscp/giwow/internal/netrc"
	"github.com/nicolasdscp/giwow/logger"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var tokenRmCmd = &cobra.Command{
	Use:   "rm [machine]",
	Short: "Remove a machine from your .netrc file",
	Args:  cobra.ExactArgs(1),
	Long:  ``,
	Run:   runTokenRm,
}

func init() {
	tokenCmd.AddCommand(tokenRmCmd)
}

func runTokenRm(_ *cobra.Command, args []string) {
	machine := netrc.Current.Machine(args[0])
	if machine == nil {
		logger.Fatal("Machine not found")
	}

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Are you sure you want to remove %s machine", args[0]),
		IsConfirm: true,
	}

	result, err := prompt.Run()
	cobra.CheckErr(err)

	if result == "y" {
		netrc.Current.RemoveMachine(args[0])
		cobra.CheckErr(netrc.Current.Save())
		logger.Print("Machine %s removed", args[0])
	} else {
		logger.Print("Machine %s not removed", args[0])
	}
}
