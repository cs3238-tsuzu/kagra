package analyzer

import (
	"strings"

	"github.com/cs3238-tsuzu/kagra/pkg/rbactypes"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
)

// ListPermissionsForRole returns associations related to all APIResources for Role
func (a *Analyzer) ListPermissionsForRole(role *rbacv1.Role) []*rbactypes.APIResourcePermissionsList {
	rpls := make([]*rbactypes.APIResourcePermissionsList, 0, len(a.apiResourceLists))
	for _, rclist := range a.apiResourceLists {
		rpl := make([]*rbactypes.APIResourcePermissions, 0, len(rclist.APIResources))

		for _, rc := range rclist.APIResources {
			perms := permissionFromPolicyRules(role.Rules, &rc, rclist.GroupVersion)

			if perms.DeniesAll() {
				continue
			}

			rpl = append(rpl, &rbactypes.APIResourcePermissions{
				APIResource: &rc,
				Permissions: perms,
			})
		}

		rpls = append(rpls, &rbactypes.APIResourcePermissionsList{
			GroupVersion: rclist.GroupVersion,
			APIResources: rpl,
		})
	}

	return rpls
}

func permissionFromPolicyRules(rules []rbacv1.PolicyRule, rc *metav1.APIResource, groupVersion string) rbactypes.Permissions {
	perms := make(rbactypes.Permissions, 0, len(rules))
	commonGroup := groupFromGroupVersion(groupVersion)

	for _, rule := range rules {
		group := commonGroup
		if rc.Group != "" {
			group = rc.Group
		}

		ok := rbacv1helpers.APIGroupMatches(&rule, group) &&
			rbacv1helpers.ResourceMatches(&rule, rc.Name, subresourceFromCombined(rc.Name))

		if !ok {
			continue
		}

		perms = append(perms, rbactypes.NewPermissionFromPolicyRule(&rule))
	}

	perms.Minify()

	return perms
}

func subresourceFromCombined(combined string) string {
	split := strings.SplitN(combined, "/", 2)

	if len(split) == 2 {
		return split[1]
	}

	return ""
}

func groupFromGroupVersion(gv string) string {
	idx := strings.Index(gv, "/")

	if idx == -1 {
		return ""
	}

	return gv[:idx]
}
