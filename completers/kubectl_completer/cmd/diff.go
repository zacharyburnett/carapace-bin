package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:     "diff -f FILENAME",
	Short:   "Diff the live version against a would-be applied version",
	GroupID: "advanced",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	diffCmd.Flags().String("field-manager", "kubectl-client-side-apply", "Name of the manager used to track field ownership.")
	diffCmd.Flags().StringSliceP("filename", "f", []string{}, "Filename, directory, or URL to files contains the configuration to diff")
	diffCmd.Flags().Bool("force-conflicts", false, "If true, server-side apply will force the changes against conflicts.")
	diffCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	diffCmd.Flags().Bool("prune", false, "Include resources that would be deleted by pruning. Can be used with -l and default shows all resources would be pruned")
	diffCmd.Flags().StringArray("prune-allowlist", []string{}, "Overwrite the default whitelist with <group/version/kind> for --prune")
	diffCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	diffCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.")
	diffCmd.Flags().Bool("server-side", false, "If true, apply runs in the server instead of the client.")
	diffCmd.Flags().Bool("show-managed-fields", false, "If true, include managed fields in the diff.")
	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
	})
}
