package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "Stop (deactivate) one or more units",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stopCmd).Standalone()

	rootCmd.AddCommand(stopCmd)

	carapace.Gen(stopCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits(stopCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
