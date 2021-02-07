package analyzer

import (
	rbacv1 "k8s.io/api/rbac/v1"
)

// GetRoleForRoleBinding returns the role specified by RoleBinding, or nil if not exists
func (a *Analyzer) GetRoleForRoleBinding(rb *rbacv1.RoleBinding) *rbacv1.Role {
	return a.GetRoleForRoleRef(rb.Namespace, &rb.RoleRef)
}

// GetRoleForRoleRef returns the role specified by RoleRef, or nil if not exists
func (a *Analyzer) GetRoleForRoleRef(namespace string, roleRef *rbacv1.RoleRef) *rbacv1.Role {
	if roleRef.APIGroup != rbacv1.GroupName {
		return nil
	}

	if roleRef.Kind != "Role" {
		return nil
	}

	roles, ok := a.opt.roles[namespace]

	if !ok {
		return nil
	}

	role, ok := roles[roleRef.Name]

	if !ok {
		return nil
	}

	return role
}

// GetClusterRoleForRoleBinding returns the role specified by RoleBinding, or nil if not exists
func (a *Analyzer) GetClusterRoleForRoleBinding(rb *rbacv1.RoleBinding) *rbacv1.ClusterRole {
	return a.GetClusterRoleForRoleRef(&rb.RoleRef)
}

// GetClusterRoleForClusterRoleBinding returns the role specified by ClusterRoleBinding, or nil if not exists
func (a *Analyzer) GetClusterRoleForClusterRoleBinding(rb *rbacv1.ClusterRoleBinding) *rbacv1.ClusterRole {
	return a.GetClusterRoleForRoleRef(&rb.RoleRef)
}

// GetClusterRoleForRoleRef returns the role specified by RoleRef, or nil if not exists
func (a *Analyzer) GetClusterRoleForRoleRef(roleRef *rbacv1.RoleRef) *rbacv1.ClusterRole {
	if roleRef.APIGroup != rbacv1.GroupName {
		return nil
	}

	if roleRef.Kind != "ClusterRole" {
		return nil
	}

	role, ok := a.opt.clusterRoles[roleRef.Name]

	if !ok {
		return nil
	}

	return role
}
