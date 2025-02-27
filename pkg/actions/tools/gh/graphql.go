package gh

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
)

func graphQlAction(opts RepoOpts, query string, v interface{}, transform func() carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		params := make([]string, 0)
		if strings.Contains(query, "$owner") {
			params = append(params, "$owner: String!")
		}
		if strings.Contains(query, "$repo") {
			params = append(params, "$repo: String!")
		}
		queryParams := strings.Join(params, ",")
		if queryParams != "" {
			queryParams = "(" + queryParams + ")"
		}

		if opts.Host == "" {
			opts.Host = "github.com"
		}

		if opts.Owner == "@me" {
			var err error
			if opts.Owner, err = userFor(opts.Host); err != nil {
				return carapace.ActionMessage(err.Error())
			}
		}

		return carapace.ActionExecCommand("gh", "api", "--hostname", opts.Host, "--header", "Accept: application/vnd.github.merge-info-preview+json", "graphql", "-F", "owner="+opts.Owner, "-F", "repo="+opts.Name, "-f", fmt.Sprintf("query=query%v {%v}", queryParams, query))(func(output []byte) carapace.Action {
			if err := json.Unmarshal(output, &v); err != nil {
				return carapace.ActionMessage("failed to unmarshall response: " + err.Error())
			}
			return transform()
		})
	})
}
