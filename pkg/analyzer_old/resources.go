package analyzer

import (
	"fmt"
	"strings"
)

// Node represents a Kubernetes resource in "rbac.authorization.k8s.io"
type Node struct {
	APIGroup string
	Kind     string

	Name      string
	Namespace string
}

func ParseNode(id string) Node {
	arr := strings.SplitN(id, "/", 5)

	node := Node{}
	switch len(arr) {
	case 4:
		node.Name = arr[4]
		fallthrough
	case 3:
		node.Namespace = arr[3]
		fallthrough
	case 2:
		node.Kind = arr[2]
		fallthrough
	case 1:
		node.APIGroup = arr[0]
	}

	return node
}

// GetUniqueID returns the unique name for the resource
func (n *Node) GetUniqueID() string {
	return fmt.Sprintf("%s/%s/%s/%s", n.APIGroup, n.Kind, n.Namespace, n.Name)
}

// Subject are applied some Roles by RoleBinding to.
type Subject struct {
	Node

	RoleBindings []string
}

func (s *Subject) SetDefault() {
	if s.Node.APIGroup == "" {
		s.Node.APIGroup = "rbac.authorization.k8s.io"
	}
}

type RoleBinding struct {
	Node

	Subjects []string
	Role     string
}

type Role struct {
	Node

	Rules []Rule
}

type Rule struct {
	Resource        string
	ResourceNames   []string
	NonResourceURLs []string

	Verbs []string
}

type Resource struct {
	Node

	Namespaced *bool
}

func NewResource(group, kind string) *Resource {
	return &Resource{
		Node: Node{
			APIGroup: group,
			Kind:     kind,
		},
	}
}
