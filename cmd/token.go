package cmd

import (
	"github.com/nicolasdscp/giwow/internal/netrc"
	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:               "token",
	Short:             "Manage your platform tokens. This will interact with your $HOME/.netrc file",
	Long:              ``,
	PersistentPreRunE: persistentPreRunEToken,
	Run:               runToken,
}

func init() {
	rootCmd.AddCommand(tokenCmd)
	tokenCmd.DisableFlagsInUseLine = true

	tokenCmd.PersistentFlags().String("netrc", "", `The path to the netrc file, default is $HOME/.netrc`)
}

func persistentPreRunEToken(cmd *cobra.Command, _ []string) error {
	return netrc.ResolveCurrent(cmd.Flag("netrc").Value.String())
}

func runToken(cmd *cobra.Command, args []string) {
	cmd.Usage()
}
