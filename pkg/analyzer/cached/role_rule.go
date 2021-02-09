package cachedanalyzer

import (
	"github.com/cs3238-tsuzu/kagra/pkg/rbactypes"
	rbacv1 "k8s.io/api/rbac/v1"
)

// Role -> Resource
// ListPermissionsForRole returns associations related to all APIResources for Role
func (a *cachedAnalyzer) ListPermissionsForRole(role *rbacv1.Role) []*rbactypes.APIResourcePermissionsList {
	key := keyForCache("ListPermissionsForRole", role)

	var res []*rbactypes.APIResourcePermissionsList
	r, hit := a.cache[key]
	if hit {
		res = r.([]*rbactypes.APIResourcePermissionsList)
	} else {
		res = a.internal.ListPermissionsForRole(role)
		a.cache[key] = res
	}

	var copied []*rbactypes.APIResourcePermissionsList
	copyViaJSON(res, &copied)

	return copied
}

func (a *cachedAnalyzer) ListPermissionsForClusterRole(role *rbacv1.ClusterRole) []*rbactypes.APIResourcePermissionsList {
	key := keyForCache("ListPermissionsForClusterRole", role)

	var res []*rbactypes.APIResourcePermissionsList
	r, hit := a.cache[key]
	if hit {
		res = r.([]*rbactypes.APIResourcePermissionsList)
	} else {
		res = a.internal.ListPermissionsForClusterRole(role)
		a.cache[key] = res
	}

	var copied []*rbactypes.APIResourcePermissionsList
	copyViaJSON(res, &copied)

	return copied
}
