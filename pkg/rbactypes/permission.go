package rbactypes

import (
	"github.com/cs3238-tsuzu/kagra/pkg/util"
	rbacv1 "k8s.io/api/rbac/v1"
)

// Permission represents a relationship b/w PolicyRule and APIResource
type Permission struct {
	Verbs         util.StringSet
	ResourceNames util.StringSet
}

// IsAllowed returns if the verb and resourceName was allowed in this permission
func (p *Permission) IsAllowed(verb, resourceName string) bool {
	if !p.IsAllowedVerb(verb) {
		return false
	}

	if len(p.ResourceNames) == 0 {
		return true
	}

	return p.ResourceNames.Contains(resourceName)
}

// IsAllowedVerb returns if the verb was allowed, but resource names are ignored.
// You should use IsAllowed instead.
func (p *Permission) IsAllowedVerb(verb string) bool {
	return p.Verbs.Contains(rbacv1.VerbAll) || p.Verbs.Contains(verb)
}

// Minify organizes the Permission
func (p *Permission) Minify() {
	if p.Verbs.Contains(rbacv1.VerbAll) {
		p.Verbs = util.NewStringSetFromSlice([]string{rbacv1.VerbAll})
	}
}

// IsAllAllowed returns if all requests are permitted by this Permission
func (p *Permission) IsAllAllowed() bool {
	return len(p.ResourceNames) == 0 && p.Verbs.Contains(rbacv1.VerbAll)
}

// Contains returns if c is covered by this Permission
func (p *Permission) Contains(c *Permission) bool {
	if len(p.ResourceNames) == 0 {
		if p.IsAllAllowed() {
			return true
		}

		return p.Verbs.ContainsSet(c.Verbs)
	}

	if !p.ResourceNames.ContainsSet(c.ResourceNames) {
		return false
	}

	return p.Verbs.ContainsSet(c.Verbs)
}
