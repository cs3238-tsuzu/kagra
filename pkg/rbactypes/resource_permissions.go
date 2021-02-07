package rbactypes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// APIResourcePermissions represents the permissions for APIResource
type APIResourcePermissions struct {
	APIResource *metav1.APIResource

	Permissions
}

// APIResourcePermissionsList the list of ResourcePermissions
type APIResourcePermissionsList struct {
	// GroupVersion is the group and version this APIResourceList is for.
	GroupVersion string

	// APIResources is the list of APIResourcePermissions
	APIResources []*APIResourcePermissions
}
