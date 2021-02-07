package analyzer

import (
	"fmt"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func apiResourceKey(rc *metav1.APIResource) string {
	return fmt.Sprintf("%s/%s", rc.Group, rc.Name)
}

func subjectKey(sbj *rbacv1.Subject) string {
	return sbj.String()
}
