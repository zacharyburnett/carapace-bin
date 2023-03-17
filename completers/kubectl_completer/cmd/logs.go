package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs [-f] [-p] (POD | TYPE/NAME) [-c CONTAINER]",
	Short: "Print the logs for a container in a pod",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logsCmd).Standalone()

	logsCmd.Flags().Bool("all-containers", false, "Get all containers' logs in the pod(s).")
	logsCmd.Flags().StringP("container", "c", "", "Print the logs of this container")
	logsCmd.Flags().BoolP("follow", "f", false, "Specify if the logs should be streamed.")
	logsCmd.Flags().Bool("ignore-errors", false, "If watching / following pod logs, allow for any errors that occur to be non-fatal")
	logsCmd.Flags().Bool("insecure-skip-tls-verify-backend", false, "Skip verifying the identity of the kubelet that logs are requested from.  In theory, an attacker could provide invalid log content back. You might want to use this if your kubelet serving certificates have expired.")
	logsCmd.Flags().Int64("limit-bytes", 0, "Maximum bytes of logs to return. Defaults to no limit.")
	logsCmd.Flags().Int("max-log-requests", 5, "Specify maximum number of concurrent logs to follow when using by a selector. Defaults to 5.")
	logsCmd.Flags().Duration("pod-running-timeout", 0, "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running")
	logsCmd.Flags().Bool("prefix", false, "Prefix each log line with the log source (pod name and container name)")
	logsCmd.Flags().BoolP("previous", "p", false, "If true, print the logs for the previous instance of the container in a pod if it exists.")
	logsCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.")
	logsCmd.Flags().Duration("since", 0, "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used.")
	logsCmd.Flags().String("since-time", "", "Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used.")
	logsCmd.Flags().Int64("tail", -1, "Lines of recent log file to display. Defaults to -1 with no selector, showing all log lines otherwise 10, if a selector is provided.")
	logsCmd.Flags().Bool("timestamps", false, "Include timestamps on each line in the log output")
	rootCmd.AddCommand(logsCmd)

	carapace.Gen(logsCmd).FlagCompletion(carapace.ActionMap{
		"container": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionMessage("no resource specified")
			} else {
				return kubectl.ActionContainers(kubectl.ContainerOpts{Namespace: "", Resource: c.Args[0]})
			}
		}),
	})

	carapace.Gen(logsCmd).PositionalCompletion(
		kubectl.ActionApiResourceResources(),
	)
}
