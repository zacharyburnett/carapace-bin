package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push [OPTIONS] [SERVICE...]",
	Short: "Push service images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	pushCmd.Flags().Bool("ignore-push-failures", false, "Push what it can and ignores images with push failures")
	pushCmd.Flags().Bool("include-deps", false, "Also push images of services declared as dependencies")
	pushCmd.Flags().BoolP("quiet", "q", false, "Push without printing progress information")
	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionServices(pushCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
