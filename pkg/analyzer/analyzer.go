package analyzer

import (
	rbacv1 "k8s.io/api/rbac/v1"
)

type Analyzer struct {
	*Option
}

func NewAnalyzer(opt *Option) *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) ListRolesForSubject(sbj *rbacv1.Subject) []*rbacv1.Role {

}

func (a *Analyzer) ListClusterRolesForSubject(sbj *rbacv1.Subject) []*rbacv1.ClusterRole {

}
