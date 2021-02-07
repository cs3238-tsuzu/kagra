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

// Allows returns if the verb and resourceName was allowed in this permission
func (p *Permission) Allows(verb, resourceName string) bool {
	if !p.AllowsVerb(verb) {
		return false
	}

	if len(p.ResourceNames) == 0 {
		return true
	}

	return p.ResourceNames.Contains(resourceName)
}

// AllowsVerb returns if the verb was allowed, but resource names are ignored.
// You should use IsAllowed instead.
func (p *Permission) AllowsVerb(verb string) bool {
	return p.Verbs.Contains(rbacv1.VerbAll) || p.Verbs.Contains(verb)
}

// Minify organizes the Permission
func (p *Permission) Minify() {
	if p.Verbs.Contains(rbacv1.VerbAll) {
		p.Verbs = util.NewStringSetFromSlice([]string{rbacv1.VerbAll})
	}
}

// AllowsAll returns if all requests are permitted by this Permission
func (p *Permission) AllowsAll() bool {
	return len(p.ResourceNames) == 0 && p.Verbs.Contains(rbacv1.VerbAll)
}

// DeniesAll returns if all requests are rejected by this Permission
func (p *Permission) DeniesAll() bool {
	return len(p.Verbs) == 0
}

// Contains returns if c is covered by this Permission
func (p *Permission) Contains(c *Permission) bool {
	if len(p.ResourceNames) == 0 {
		if p.AllowsAll() {
			return true
		}

		return p.Verbs.ContainsSet(c.Verbs)
	}

	if !p.ResourceNames.ContainsSet(c.ResourceNames) {
		return false
	}

	return p.Verbs.Contains(rbacv1.VerbAll) || p.Verbs.ContainsSet(c.Verbs)
}
