package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cache_deleteCmd = &cobra.Command{
	Use:   "delete [<cache-id>| <cache-key> | --all]",
	Short: "Delete Github Actions caches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_deleteCmd).Standalone()

	cache_deleteCmd.Flags().BoolP("all", "a", false, "Delete all caches")
	cacheCmd.AddCommand(cache_deleteCmd)

	carapace.Gen(cache_deleteCmd).PositionalCompletion(
		action.ActionCaches(cache_deleteCmd),
	)
}
