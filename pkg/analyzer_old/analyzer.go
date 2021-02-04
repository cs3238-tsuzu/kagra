package analyzer

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Analyzer struct {
	resources    map[string]*Resource
	roles        map[string]*Role
	roleBindings map[string]*RoleBinding
}

// NewAnalyzer initializes a new analyzer for RBAC resources
func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) addResouce(r *Resource) {
	a.resources[r.GetUniqueID()] = r
}

func (a *Analyzer) AddAPIResource(resource *metav1.APIResource) {
	rc := NewResource(resource.Group, resource.Kind)
	rc.Namespaced = &resource.Namespaced

	a.addResouce(rc)
}

func (a *Analyzer) AddRole(role *rbacv1.Role) {
	group, _ := parseAPIVersino(role.APIVersion)

	rules := make([]Rule, 0, len(role.Rules)*3)

	for _, rule := range role.Rules {
		for _, g := range rule.APIGroups {
			for _, r := range rule.Resources {
				rc := NewResource(g, r)

				a.addResouce(rc)

				rules = append(rules, Rule{
					Resource:        rc.GetUniqueID(),
					ResourceNames:   rule.ResourceNames,
					NonResourceURLs: rule.NonResourceURLs,
					Verbs:           rule.Verbs,
				})
			}
		}
	}

	r := &Role{
		Node: Node{
			APIGroup: group,
			Kind:     role.Kind,

			Namespace: role.Namespace,
			Name:      role.Name,
		},
		Rules: rules,
	}

	a.roles[r.GetUniqueID()] = r
}

func (a *Analyzer) AddRoleBinding(roleBinding *rbacv1.RoleBinding) {
	group, _ := parseAPIVersino(roleBinding.APIVersion)

	roleID := a.parseRoleRef(roleBinding.Namespace, roleBinding.RoleRef)

	subjects := make([]string, 0, len(roleBinding.Subjects))
	for _, subject := range roleBinding.Subjects {
		sbj := Subject{
			Node: Node{
				APIGroup:  subject.APIGroup,
				Kind:      subject.Kind,
				Name:      subject.Name,
				Namespace: subject.Namespace,
			},
		}

		subjects = append(subjects, sbj.GetUniqueID())
	}

	rb := RoleBinding{
		Subjects: subjects,
		Role:     roleID,
	}

}
