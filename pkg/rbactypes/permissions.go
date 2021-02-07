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

// DeniesAll returns if all requests will be rejected by these permissions
func (pms Permissions) DeniesAll() bool {
	for _, p := range pms {
		if !p.DeniesAll() {
			return false
		}
	}

	return true
}

// Filter filters Permissions by walk function
func (pms Permissions) Filter(walk PermissionsFilter) Permissions {
	res := make(Permissions, 0)

	for _, p := range pms {
		if walk(p) {
			res = append(res, p)
		}
	}

	return res
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
