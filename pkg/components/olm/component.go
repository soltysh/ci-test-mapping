package olm

import (
	v1 "github.com/openshift-eng/ci-test-mapping/pkg/api/types/v1"
	"github.com/openshift-eng/ci-test-mapping/pkg/config"
	"github.com/openshift-eng/ci-test-mapping/pkg/util"
)

type Component struct {
	*config.Component
}

var OLMComponent = Component{
	Component: &config.Component{
		Name:                 "OLM",
		Operators:            []string{"marketplace", "operator-lifecycle-manager", "operator-lifecycle-manager-catalog", "operator-lifecycle-manager-packageserver"},
		DefaultJiraComponent: "OLM",
		Matchers: []config.ComponentMatcher{
			{
				IncludeSubstrings: []string{"bz-OLM"},
			},
			{
				IncludeSubstrings: []string{"bz-platform-operators-aggregated"},
			},
			{
				SIG: "sig-operator",
			},
		},
	},
}

func (c *Component) IdentifyTest(test *v1.TestInfo) (*v1.TestOwnership, error) {
	if matcher := c.FindMatch(test); matcher != nil {
		return &v1.TestOwnership{
			Name:          test.Name,
			Component:     c.Name,
			JIRAComponent: matcher.JiraComponent,
			Priority:      matcher.Priority,
			Capabilities:  append(matcher.Capabilities, identifyCapabilities(test)...),
		}, nil
	}

	return nil, nil
}

func (c *Component) StableID(test *v1.TestInfo) string {
	return util.StableID(test, nil)
}

func (c *Component) JiraComponents() (components []string) {
	components = []string{c.DefaultJiraComponent}
	for _, m := range c.Matchers {
		components = append(components, m.JiraComponent)
	}

	return components
}
