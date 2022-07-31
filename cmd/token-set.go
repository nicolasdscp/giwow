package cmd

import (
	"github.com/nicolasdscp/giwow/internal/netrc"
	"github.com/nicolasdscp/giwow/logger"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var tokenSetCmd = &cobra.Command{
	Use:   "set [machine]",
	Short: "Set values for a machine",
	Args:  cobra.ExactArgs(1),
	Long:  ``,
	Run:   runTokenSet,
}

func init() {
	tokenCmd.AddCommand(tokenSetCmd)

	tokenSetCmd.Flags().StringP("user", "u", "", "Set the user")
	tokenSetCmd.Flags().StringP("password", "p", "", "Set the password")
}

func runTokenSet(cmd *cobra.Command, args []string) {
	user := cmd.Flag("user").Value.String()
	password := cmd.Flag("password").Value.String()
	machine := netrc.Current.Machine(args[0])

	if user == "" && password == "" {
		logger.Fatal("No value to set")
	}

	if machine == nil {
		logger.Fatal("Machine not found")
	}

	if user != "" {
		machine.Set("user", user)
	}

	if password != "" {
		machine.Set("password", password)
	}

	cobra.CheckErr(netrc.Current.Save())

	logger.Print("Machine %s updated", args[0])
}
