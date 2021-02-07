package rbactypes

import (
	"testing"

	"github.com/cs3238-tsuzu/kagra/pkg/util"
	"github.com/google/go-cmp/cmp"
	rbacv1 "k8s.io/api/rbac/v1"
)

func TestPermissions_Minify(t *testing.T) {
	tests := []struct {
		name     string
		pms      Permissions
		expected Permissions
	}{
		{
			name: "generally allowed",
			pms: Permissions([]*Permission{
				{
					Verbs:         util.NewStringSet("create"),
					ResourceNames: util.NewStringSet("a", "b", "c"),
				},
				{
					Verbs: util.NewStringSet(rbacv1.VerbAll),
				},
			}),
			expected: Permissions([]*Permission{
				{
					Verbs: util.NewStringSet(rbacv1.VerbAll),
				},
			}),
		},
		{
			name: "omitted due to already contained by general one",
			pms: Permissions([]*Permission{
				{
					Verbs:         util.NewStringSet("create"),
					ResourceNames: util.NewStringSet("a", "b", "c"),
				},
				{
					Verbs: util.NewStringSet("create", "update"),
				},
			}),
			expected: Permissions([]*Permission{
				{
					Verbs: util.NewStringSet("create", "update"),
				},
			}),
		},
		{
			name: "omitted due to already contained by previous one",
			pms: Permissions([]*Permission{
				{
					Verbs:         util.NewStringSet("create", "update"),
					ResourceNames: util.NewStringSet("a", "b", "c"),
				},
				{
					Verbs:         util.NewStringSet("create"),
					ResourceNames: util.NewStringSet("a", "b"),
				},
			}),
			expected: Permissions([]*Permission{
				{
					Verbs:         util.NewStringSet("create", "update"),
					ResourceNames: util.NewStringSet("a", "b", "c"),
				},
			}),
		},
		{
			name: "combined",
			pms: Permissions([]*Permission{
				{
					Verbs: util.NewStringSet("create", "update"),
				},
				{
					Verbs: util.NewStringSet("get"),
				},
			}),
			expected: Permissions([]*Permission{
				{
					Verbs: util.NewStringSet("create", "update", "get"),
				},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pms.Minify()

			if diff := cmp.Diff(tt.pms, tt.expected); diff != "" {
				t.Errorf("Permission.Minify(): %s", diff)
			}
		})
	}
}
