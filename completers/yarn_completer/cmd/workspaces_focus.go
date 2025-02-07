package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/yarn"
	"github.com/spf13/cobra"
)

var workspaces_focusCmd = &cobra.Command{
	Use:   "focus",
	Short: "Install a single workspace and its dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_focusCmd).Standalone()

	workspaces_focusCmd.Flags().BoolP("all", "A", false, "Install the entire project")
	workspaces_focusCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	workspaces_focusCmd.Flags().Bool("production", false, "Only install regular dependencies by omitting dev dependencies")
	workspacesCmd.AddCommand(workspaces_focusCmd)

	carapace.Gen(workspaces_focusCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return yarn.ActionWorkspaces().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
