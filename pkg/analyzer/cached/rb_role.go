package cachedanalyzer

import (
	rbacv1 "k8s.io/api/rbac/v1"
)

// RoleBinding -> Role
// The cache won't be efficient

// GetRoleForRoleBinding returns the role specified by RoleBinding, or nil if not exists
func (a *cachedAnalyzer) GetRoleForRoleBinding(rb *rbacv1.RoleBinding) *rbacv1.Role {
	return a.internal.GetRoleForRoleBinding(rb)
}

// GetRoleForRoleRef returns the role specified by RoleRef, or nil if not exists
func (a *cachedAnalyzer) GetRoleForRoleRef(namespace string, roleRef *rbacv1.RoleRef) *rbacv1.Role {
	return a.internal.GetRoleForRoleRef(namespace, roleRef)
}

// GetClusterRoleForRoleBinding returns the role specified by RoleBinding, or nil if not exists
func (a *cachedAnalyzer) GetClusterRoleForRoleBinding(rb *rbacv1.RoleBinding) *rbacv1.ClusterRole {
	return a.internal.GetClusterRoleForRoleBinding(rb)
}

// GetClusterRoleForClusterRoleBinding returns the role specified by ClusterRoleBinding, or nil if not exists
func (a *cachedAnalyzer) GetClusterRoleForClusterRoleBinding(rb *rbacv1.ClusterRoleBinding) *rbacv1.ClusterRole {
	return a.internal.GetClusterRoleForClusterRoleBinding(rb)
}

// GetClusterRoleForRoleRef returns the role specified by RoleRef, or nil if not exists
func (a *cachedAnalyzer) GetClusterRoleForRoleRef(roleRef *rbacv1.RoleRef) *rbacv1.ClusterRole {
	return a.internal.GetClusterRoleForRoleRef(roleRef)
}
