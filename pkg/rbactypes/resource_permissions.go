package rbactypes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// APIResourcePermissions represents the permissions for APIResource
type APIResourcePermissions struct {
	APIResource *metav1.APIResource

	Permissions
}

// ResourcePermissionsList the list of ResourcePermissions
type ResourcePermissionsList struct {
	// GroupVersion is the group and version this APIResourceList is for.
	GroupVersion string

	// APIResources list ups
	APIResources []APIResourcePermissions
}
