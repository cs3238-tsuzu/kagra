package rbactypes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ResourcePermissions represents the permissions for APIResource
type ResourcePermissions struct {
	Resource *metav1.APIResource

	Permissions
}
