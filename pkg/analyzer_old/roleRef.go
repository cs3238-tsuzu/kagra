package analyzer

import (
	"strings"

	rbacv1 "k8s.io/api/rbac/v1"
)

func (a *Analyzer) parseRoleRef(namespace string, roleRef rbacv1.RoleRef) string {
	role := Role{
		Node: Node{
			APIGroup: roleRef.APIGroup,
			Kind:     roleRef.Kind,
			Name:     roleRef.Name,
		},
	}

	namespaced := a.isNamespacedRoleRef(roleRef)

	if namespaced {
		role.Namespace = namespace
	}

	return role.GetUniqueID()
}

func (a *Analyzer) isNamespacedRoleRef(roleRef rbacv1.RoleRef) bool {
	resourceID := NewResource(roleRef.APIGroup, roleRef.Kind).GetUniqueID()

	rc, ok := a.resources[resourceID]

	if ok && rc.Namespaced != nil {
		return *rc.Namespaced
	}

	return !strings.HasPrefix(roleRef.Kind, "Cluster")
}
