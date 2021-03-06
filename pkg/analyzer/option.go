package analyzer

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Option has all resources for Analyzer
type Option struct {
	apiResourceLists    map[string]*metav1.APIResourceList
	roles               map[string]map[string]*rbacv1.Role
	roleBindings        map[string]map[string]*rbacv1.RoleBinding
	clusterRoles        map[string]*rbacv1.ClusterRole
	clusterRoleBindings map[string]*rbacv1.ClusterRoleBinding
	subjects            map[string]*rbacv1.Subject
}

// NewOption initializes a new empty Option
func NewOption() *Option {
	return &Option{}
}

// AddAPIResourceList adds a new APIResourceList
func (o *Option) AddAPIResourceList(apiResourceList *metav1.APIResourceList) *Option {
	if o.apiResourceLists == nil {
		o.apiResourceLists = make(map[string]*metav1.APIResourceList)
	}

	o.apiResourceLists[apiResourceList.GroupVersion] = apiResourceList

	return o
}

// AddRole adds a new Role
func (o *Option) AddRole(role *rbacv1.Role) {
	if o.roles == nil {
		o.roles = make(map[string]map[string]*rbacv1.Role)
	}

	if o.roles[role.Namespace] == nil {
		o.roles[role.Namespace] = make(map[string]*rbacv1.Role)
	}

	o.roles[role.Namespace][role.Name] = role
}

// AddRoleBinding adds a new RoleBinding
func (o *Option) AddRoleBinding(rb *rbacv1.RoleBinding) {
	if o.roleBindings == nil {
		o.roleBindings = make(map[string]map[string]*rbacv1.RoleBinding)
	}

	if o.roleBindings[rb.Namespace] == nil {
		o.roleBindings[rb.Namespace] = make(map[string]*rbacv1.RoleBinding)
	}

	o.roleBindings[rb.Namespace][rb.Name] = rb
}

// AddSubject adds a new Subject
func (o *Option) AddSubject(sbj *rbacv1.Subject) {
	if o.subjects == nil {
		o.subjects = make(map[string]*rbacv1.Subject)
	}

	o.subjects[subjectKey(sbj)] = sbj
}
