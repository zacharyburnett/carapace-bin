package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dlxCmd = &cobra.Command{
	Use:     "dlx",
	Short:   "Run a package in a temporary environment",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dlxCmd).Standalone()
	dlxCmd.Flags().SetInterspersed(false)

	dlxCmd.Flags().StringSliceP("package", "p", []string{}, "The package(s) to install before running the command")
	dlxCmd.Flags().BoolP("quiet", "q", false, "Only report critical errors instead of printing the full install logs")
	rootCmd.AddCommand(dlxCmd)

	// TODO package completion

	carapace.Gen(dlxCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	)

	carapace.Gen(dlxCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return bridge.ActionCarapaceBin(c.Args[0]).Shift(1)
		}),
	)
}
