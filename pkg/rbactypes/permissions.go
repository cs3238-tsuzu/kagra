package rbactypes

import (
	"github.com/cs3238-tsuzu/kagra/pkg/util"
)

// Permissions bundles some Permission reousrces
type Permissions []*Permission

// Allows returns if the verb and resourceName was allowed in these permissions
func (pms Permissions) Allows(verb, resourceName string) bool {
	for _, p := range pms {
		if p.Allows(verb, resourceName) {
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
		specific = append(specific, pm)
	}

	general := &Permission{
		Verbs: generalVerbs,
	}

	if general.AllowsAll() {
		*pms = []*Permission{general}

		return
	}

	*pms = make(Permissions, 0, len(specific)+1)

	if len(general.Verbs) != 0 {
		*pms = append(*pms, general)
	}

	for i := range specific {
		specific[i].Minify()

		if general.Contains(specific[i]) {
			continue
		}

		skipped := false
		for j := 0; j < i; j++ {
			if specific[j].Contains(specific[i]) {
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
