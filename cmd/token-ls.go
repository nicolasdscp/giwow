package cmd

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/nicolasdscp/giwow/internal/netrc"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var tokenLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List your auth tokens",
	Long:  `This will basically list your tokens in your $HOME/.netrc file.`,
	Run:   runTokenLs,
}

func init() {
	tokenCmd.AddCommand(tokenLsCmd)

	tokenLsCmd.Flags().BoolP("magnify", "m", false, "Print tokens in a magnified array format")
	tokenLsCmd.Flags().Bool("showPass", false, "Display your tokens passwords")
}

func runTokenLs(cmd *cobra.Command, _ []string) {
	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)

	if !(cmd.Flag("magnify").Value.String() == "true") {
		t.Style().Options.SeparateColumns = false
		t.Style().Options.DrawBorder = false
		t.Style().Options.SeparateHeader = true
	}

	t.AppendHeader(table.Row{"#", "Machine", "Login", "Password"})
	machines := netrc.Current.GetMachines()
	var pass string
	for i, m := range machines {
		pass = "********"

		if cmd.Flag("showPass").Value.String() == "true" {
			pass = m.Get("password")
		}

		if m.IsDefault {
			t.AppendRow(table.Row{"D", m.Name, m.Get("login"), pass})
			break
		}
		t.AppendRow(table.Row{i + 1, m.Name, m.Get("login"), pass})
	}
	t.Render()
}
