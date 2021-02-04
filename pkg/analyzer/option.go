package analyzer

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Option struct {
	apiResources        map[string]*metav1.APIResource
	roles               map[string]map[string]*rbacv1.Role
	roleBindings        map[string]map[string]*rbacv1.RoleBinding
	clusterRoles        map[string]*rbacv1.ClusterRole
	clusterRoleBindings map[string]*rbacv1.ClusterRoleBinding
	subjects            map[string]*rbacv1.Subject
}

func NewOption() *Option {
	return &Option{}
}

func (o *Option) AddAPIResource(apiResource *metav1.APIResource) *Option {
	if o.apiResources == nil {
		o.apiResources = make(map[string]*metav1.APIResource)
	}

	o.apiResources[apiResourceKey(apiResource)] = apiResource

	return o
}

func (o *Option) AddRole(role *rbacv1.Role) {
	if o.roles == nil {
		o.roles = make(map[string]map[string]*rbacv1.Role)
	}

	if o.roles[role.Namespace] == nil {
		o.roles[role.Namespace] = make(map[string]*rbacv1.Role)
	}

	o.roles[role.Namespace][role.Name] = role
}

func (o *Option) AddRoleBinding(rb *rbacv1.RoleBinding) {
	if o.roleBindings == nil {
		o.roleBindings = make(map[string]map[string]*rbacv1.RoleBinding)
	}

	if o.roleBindings[rb.Namespace] == nil {
		o.roleBindings[rb.Namespace] = make(map[string]*rbacv1.RoleBinding)
	}

	o.roleBindings[rb.Namespace][rb.Name] = rb
}

func (o *Option) AddSubject(sbj *rbacv1.Subject) {
	if o.subjects == nil {
		o.subjects = make(map[string]*rbacv1.Subject)
	}

	o.subjects[subjectKey(sbj)] = sbj
}
