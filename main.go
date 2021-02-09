package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/yaml"

	//...
	// "k8s.io/client-go/pkg/api/v1"
	// "k8s.io/client-go/pkg/apis/rbac/v1beta1"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// const yaml = `
// apiVersion: rbac.authorization.k8s.io/v1
// kind: Role
// metadata:
//   name: hoge
// rules:
// - apiGroups: ["apps"]
// `

func main() {
	// deserializer := scheme.Codecs.UniversalDeserializer()

	// obj, _, err := deserializer.Decode([]byte(yaml), nil, nil)

	// if err != nil {
	// 	panic(err)
	// }

	// role := obj.(*rbacv1.Role)

	// fmt.Println(role)

	// viz := graphviz.New()
	// defer viz.Close()

	// g, err := viz.Graph()

	// if err != nil {
	// 	panic(err)
	// }
	// defer g.Close()

	// node, err := g.CreateNode("foo")

	// if err != nil {
	// 	panic(err)
	// }

	// node.SetHref("foobar%%%%%%%%%%%")

	// var buf bytes.Buffer
	// viz.Render(g, graphviz.XDOT, &buf)
	// fmt.Println(buf.String())

	// viz.RenderFilename(g, graphviz.PNG, "graph.png")

	var kubeconfig *string
	if home, _ := os.UserHomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// clientset, err := kubernetes.NewForConfig(config)

	// if err != nil {
	// 	panic(err)
	// }

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)

	if err != nil {
		panic(err)
	}

	rcs, err := discoveryClient.ServerResources()

	if err != nil {
		panic(err)
	}

	b, err := yaml.Marshal(rcs)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
