package rbactypes

import (
	"github.com/cs3238-tsuzu/kagra/pkg/util"
)

// Permissions bundles some Permission reousrces
type Permissions []*Permission

// IsAllowed returns if the verb and resourceName was allowed in these permissions
func (pms Permissions) IsAllowed(verb, resourceName string) bool {
	for _, p := range pms {
		if p.IsAllowed(verb, resourceName) {
			return true
		}
	}

	return false
}

// Minify organizes permissions
// TODO: Better algorithm (current: O(N^2))
func (pms *Permissions) Minify() {
	generalVerbs := util.NewStringSet()
	specific := make(Permissions, 0, len(*pms))

	for _, pm := range *pms {
		if len(pm.ResourceNames) == 0 {
			for v := range pm.Verbs {
				generalVerbs.Insert(v)
			}

			continue
		}
	}

	general := &Permission{
		Verbs: generalVerbs,
	}

	if general.IsAllAllowed() {
		*pms = []*Permission{general}

		return
	}

	*pms = make(Permissions, 0, len(specific)+1)

	for i := range specific {
		specific[i].Minify()

		if general.Contains(specific[i]) {
			continue
		}

		skipped := false
		for j := 0; j < i; j++ {
			if specific[i].Contains(specific[j]) {
				skipped = true
				break
			}
		}
		if skipped {
			continue
		}

		*pms = append(*pms, specific[i])
	}
}
