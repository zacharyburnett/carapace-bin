package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var freezeCmd = &cobra.Command{
	Use:     "freeze",
	Short:   "Freeze execution of unit processes",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(freezeCmd).Standalone()

	rootCmd.AddCommand(freezeCmd)

	carapace.Gen(freezeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits(freezeCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
