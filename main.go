package main

import (
	"bytes"
	"fmt"

	"github.com/goccy/go-graphviz"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/client-go/kubernetes/scheme"
	//...
	// "k8s.io/client-go/pkg/api/v1"
	// "k8s.io/client-go/pkg/apis/rbac/v1beta1"
)

const yaml = `
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: hoge
rules:
- apiGroups: ["apps"]
`

func main() {
	deserializer := scheme.Codecs.UniversalDeserializer()

	obj, _, err := deserializer.Decode([]byte(yaml), nil, nil)

	if err != nil {
		panic(err)
	}

	role := obj.(*rbacv1.Role)

	fmt.Println(role)

	viz := graphviz.New()
	defer viz.Close()

	g, err := viz.Graph()

	if err != nil {
		panic(err)
	}
	defer g.Close()

	node, err := g.CreateNode("foo")

	if err != nil {
		panic(err)
	}

	node.SetHref("foobar%%%%%%%%%%%")

	var buf bytes.Buffer
	viz.Render(g, graphviz.XDOT, &buf)
	fmt.Println(buf.String())

	viz.RenderFilename(g, graphviz.PNG, "graph.png")
}
