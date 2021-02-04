package analyzer

import (
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apiserver/pkg/authentication/serviceaccount"
	"k8s.io/apiserver/pkg/authentication/user"
)

// ListRoleBindingsForSubject returns all role bindings for the subject in the namespace(empty for all namespaces)
func (a *Analyzer) ListRoleBindingsForSubject(sbj *rbacv1.Subject, namespace string) []*rbacv1.RoleBinding {
	u := a.getUserInfoForSubject(sbj)

	matched := make([]*rbacv1.RoleBinding, 0)

	for ns, rbs := range a.roleBindings {
		if !(ns == "" || ns == namespace) {
			continue
		}

		for _, rb := range rbs {
			_, ok := appliesTo(u, rb.Subjects, rb.Namespace)

			if ok {
				matched = append(matched, rb)
			}
		}
	}

	return matched
}

// ListClusterRoleBindingsForSubject returns all role bindings for the subject in the namespace(empty for all namespaces)
func (a *Analyzer) ListClusterRoleBindingsForSubject(sbj *rbacv1.Subject) []*rbacv1.ClusterRoleBinding {
	u := a.getUserInfoForSubject(sbj)

	matched := make([]*rbacv1.ClusterRoleBinding, 0)

	for _, rb := range a.clusterRoleBindings {
		_, ok := appliesTo(u, rb.Subjects, rb.Namespace)

		if ok {
			matched = append(matched, rb)
		}
	}

	return matched
}

func (a *Analyzer) getUserInfoForSubject(sbj *rbacv1.Subject) user.Info {
	var u *user.DefaultInfo
	switch sbj.Kind {
	case rbacv1.UserKind:
		u = &user.DefaultInfo{
			Name: sbj.Name,
		}
	case rbacv1.GroupKind:
		u = &user.DefaultInfo{
			Groups: []string{sbj.Name},
		}
	case rbacv1.ServiceAccountKind:
		u = &user.DefaultInfo{
			Name: serviceaccount.MakeUsername(sbj.Namespace, sbj.Name),
		}
	default:
		panic("unknown kind for the subject: " + sbj.Kind)
	}

	return u
}

