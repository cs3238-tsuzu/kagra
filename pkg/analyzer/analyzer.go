package analyzer

import (
	"github.com/cs3238-tsuzu/kagra/pkg/rbactypes"
	rbacv1 "k8s.io/api/rbac/v1"
)

// Analyzer analyzes Kubernetes RBAC resources
type Analyzer interface {
	// Subject -> RoleBinding

	// ListRoleBindingsForSubject returns all role bindings for the subject in the namespace(empty for all namespaces)
	ListRoleBindingsForSubject(sbj *rbacv1.Subject, namespace string) []*rbacv1.RoleBinding

	// ListClusterRoleBindingsForSubject returns all role bindings for the subject in the namespace(empty for all namespaces)
	ListClusterRoleBindingsForSubject(sbj *rbacv1.Subject) []*rbacv1.ClusterRoleBinding

	// RoleBinding -> Role

	// GetRoleForRoleBinding returns the role specified by RoleBinding, or nil if not exists
	GetRoleForRoleBinding(rb *rbacv1.RoleBinding) *rbacv1.Role

	// GetRoleForRoleRef returns the role specified by RoleRef, or nil if not exists
	GetRoleForRoleRef(namespace string, roleRef *rbacv1.RoleRef) *rbacv1.Role

	// GetClusterRoleForRoleBinding returns the role specified by RoleBinding, or nil if not exists
	GetClusterRoleForRoleBinding(rb *rbacv1.RoleBinding) *rbacv1.ClusterRole

	// GetClusterRoleForClusterRoleBinding returns the role specified by ClusterRoleBinding, or nil if not exists
	GetClusterRoleForClusterRoleBinding(rb *rbacv1.ClusterRoleBinding) *rbacv1.ClusterRole

	// GetClusterRoleForRoleRef returns the role specified by RoleRef, or nil if not exists
	GetClusterRoleForRoleRef(roleRef *rbacv1.RoleRef) *rbacv1.ClusterRole

	// Role -> Resource

	// ListPermissionsForRole returns associations related to all APIResources for Role
	ListPermissionsForRole(role *rbacv1.Role) []*rbactypes.APIResourcePermissionsList

	ListPermissionsForClusterRole(role *rbacv1.ClusterRole) []*rbactypes.APIResourcePermissionsList
}

type analyzer struct {
	opt *Option
}

// NewAnalyzer initializes an Analyzer
func NewAnalyzer(opt *Option) Analyzer {
	return &analyzer{
		opt: opt,
	}
}
