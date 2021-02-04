package analyzer

type dummyUser struct {
	name, uid string
	groups    []string
	extra     map[string][]string
}

// GetName returns the name that uniquely identifies this user among all
// other active users.
func (u *dummyUser) GetName() string {
	return u.name
}

// GetUID returns a unique value for a particular user that will change
// if the user is removed from the system and another user is added with
// the same name.
func (u *dummyUser) GetUID() string {
	return u.uid
}

// GetGroups returns the names of the groups the user is a member of
func (u *dummyUser) GetGroups() []string {
	return u.groups
}

// GetExtra can contain any additional information that the authenticator
// thought was interesting.  One example would be scopes on a token.
// Keys in this map should be namespaced to the authenticator or
// authenticator/authorizer pair making use of them.
// For instance: "example.org/foo" instead of "foo"
// This is a map[string][]string because it needs to be serializeable into
// a SubjectAccessReviewSpec.authorization.k8s.io for proper authorization
// delegation flows
// In order to faithfully round-trip through an impersonation flow, these keys
// MUST be lowercase.
func (u *dummyUser) GetExtra() map[string][]string {
	return u.extra
}
