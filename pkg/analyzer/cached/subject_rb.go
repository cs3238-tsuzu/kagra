package cachedanalyzer

import (
	rbacv1 "k8s.io/api/rbac/v1"
)

// Subject -> RoleBinding
// ListRoleBindingsForSubject returns all role bindings for the subject in the namespace(empty for all namespaces)
func (a *cachedAnalyzer) ListRoleBindingsForSubject(sbj *rbacv1.Subject, namespace string) []*rbacv1.RoleBinding {
	key := keyForCache("ListRoleBindingsForSubject", sbj, namespace)

	var res []*rbacv1.RoleBinding
	r, hit := a.cache[key]
	if hit {
		res = r.([]*rbacv1.RoleBinding)
	} else {
		res = a.internal.ListRoleBindingsForSubject(sbj, namespace)
		a.cache[key] = res
	}

	copied := make([]*rbacv1.RoleBinding, 0, len(res))
	for i := range res {
		copied = append(copied, res[i].DeepCopy())
	}

	return copied
}

// ListClusterRoleBindingsForSubject returns all role bindings for the subject in the namespace(empty for all namespaces)
func (a *cachedAnalyzer) ListClusterRoleBindingsForSubject(sbj *rbacv1.Subject) []*rbacv1.ClusterRoleBinding {
	key := keyForCache("ListClusterRoleBindingsForSubject", sbj)

	var res []*rbacv1.ClusterRoleBinding
	r, hit := a.cache[key]
	if hit {
		res = r.([]*rbacv1.ClusterRoleBinding)
	} else {
		res = a.internal.ListClusterRoleBindingsForSubject(sbj)
		a.cache[key] = res
	}

	copied := make([]*rbacv1.ClusterRoleBinding, 0, len(res))
	for i := range res {
		copied = append(copied, res[i].DeepCopy())
	}

	return copied
}
