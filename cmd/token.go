package cmd

import (
	"github.com/nicolasdscp/giwow/internal/exception"
	"github.com/nicolasdscp/giwow/internal/netrc"
	"github.com/nicolasdscp/giwow/logger"

	"os"

	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:               "token",
	Short:             "Manage your platform tokens. This will interact with your $HOME/.netrc file",
	Long:              ``,
	PersistentPreRunE: persistentPreRunEToken,
	RunE:              runTokenE,
}

func init() {
	rootCmd.AddCommand(tokenCmd)
	tokenCmd.DisableFlagsInUseLine = true

	tokenCmd.PersistentFlags().String("netrc", "", `The path to the netrc file, default is $HOME/.netrc`)
}

func persistentPreRunEToken(cmd *cobra.Command, _ []string) error {
	err := netrc.ResolveCurrent(cmd.Flag("netrc").Value.String())
	if os.IsNotExist(err) {
		logger.Fatal(exception.ErrNetrcFileNotExist(err).Error())
	}

	return err
}

func runTokenE(cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}
