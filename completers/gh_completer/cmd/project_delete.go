package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_deleteCmd = &cobra.Command{
	Use:   "delete [<number>]",
	Short: "Delete a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_deleteCmd).Standalone()

	project_deleteCmd.Flags().String("format", "", "Output format: {json}")
	project_deleteCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	projectCmd.AddCommand(project_deleteCmd)

	carapace.Gen(project_deleteCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
		"owner":  gh.ActionOwners(gh.HostOpts{}),
	})

	carapace.Gen(project_deleteCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionProjects(gh.ProjectOpts{
				Owner:  project_deleteCmd.Flag("owner").Value.String(),
				Open:   true,
				Closed: true,
			})
		}),
	)
}
