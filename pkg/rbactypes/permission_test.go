package rbactypes

import (
	"testing"

	"github.com/cs3238-tsuzu/kagra/pkg/util"
	rbacv1 "k8s.io/api/rbac/v1"
)

func TestPermission_IsAllowed(t *testing.T) {
	type fields struct {
		Verbs         util.StringSet
		ResourceNames util.StringSet
	}
	type args struct {
		verb         string
		resourceName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "normal",
			fields: fields{
				Verbs:         util.NewStringSetFromSlice([]string{"create", "update"}),
				ResourceNames: nil,
			},
			args: args{
				verb:         "create",
				resourceName: "test",
			},
			want: true,
		},
		{
			name: "verb all",
			fields: fields{
				Verbs:         util.NewStringSetFromSlice([]string{rbacv1.VerbAll}),
				ResourceNames: nil,
			},
			args: args{
				verb:         "create",
				resourceName: "test",
			},
			want: true,
		},
		{
			name: "resource name",
			fields: fields{
				Verbs:         util.NewStringSetFromSlice([]string{rbacv1.VerbAll}),
				ResourceNames: util.NewStringSetFromSlice([]string{"test", "test2"}),
			},
			args: args{
				verb:         "create",
				resourceName: "test",
			},
			want: true,
		},
		{
			name: "denied by verb",
			fields: fields{
				Verbs:         util.NewStringSetFromSlice([]string{"create"}),
				ResourceNames: util.NewStringSetFromSlice([]string{"test", "test2"}),
			},
			args: args{
				verb:         "update",
				resourceName: "test",
			},
			want: false,
		},
		{
			name: "denied by resource name",
			fields: fields{
				Verbs:         util.NewStringSetFromSlice([]string{"create"}),
				ResourceNames: util.NewStringSetFromSlice([]string{"test", "test2"}),
			},
			args: args{
				verb:         "update",
				resourceName: "test3",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Permission{
				Verbs:         tt.fields.Verbs,
				ResourceNames: tt.fields.ResourceNames,
			}
			if got := p.IsAllowed(tt.args.verb, tt.args.resourceName); got != tt.want {
				t.Errorf("Permission.IsAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermission_Contains(t *testing.T) {
	type fields struct {
		Verbs         util.StringSet
		ResourceNames util.StringSet
	}
	type args struct {
		c *Permission
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "true by verb all - 1",
			fields: fields{
				Verbs:         util.NewStringSetFromSlice([]string{rbacv1.VerbAll}),
				ResourceNames: util.NewStringSetFromSlice([]string{"test", "test2"}),
			},
			args: args{
				c: &Permission{
					Verbs:         util.NewStringSetFromSlice([]string{rbacv1.VerbAll}),
					ResourceNames: util.NewStringSetFromSlice([]string{"test"}),
				},
			},
			want: true,
		},
		{
			name: "true by verb all - 2",
			fields: fields{
				Verbs:         util.NewStringSetFromSlice([]string{rbacv1.VerbAll}),
				ResourceNames: util.NewStringSetFromSlice([]string{"test", "test2"}),
			},
			args: args{
				c: &Permission{
					Verbs:         util.NewStringSetFromSlice([]string{"create"}),
					ResourceNames: util.NewStringSetFromSlice([]string{"test"}),
				},
			},
			want: true,
		},
		{
			name: "true by empty resource names all - 1",
			fields: fields{
				Verbs:         util.NewStringSetFromSlice([]string{"create"}),
				ResourceNames: nil,
			},
			args: args{
				c: &Permission{
					Verbs:         util.NewStringSetFromSlice([]string{"create"}),
					ResourceNames: util.NewStringSetFromSlice([]string{"test"}),
				},
			},
			want: true,
		},
		{
			name: "false by verb",
			fields: fields{
				Verbs:         util.NewStringSetFromSlice([]string{"create"}),
				ResourceNames: nil,
			},
			args: args{
				c: &Permission{
					Verbs:         util.NewStringSetFromSlice([]string{"create", "update"}),
					ResourceNames: util.NewStringSetFromSlice([]string{"test"}),
				},
			},
			want: false,
		},
		{
			name: "false by resource names",
			fields: fields{
				Verbs:         util.NewStringSetFromSlice([]string{"create"}),
				ResourceNames: util.NewStringSetFromSlice([]string{"test"}),
			},
			args: args{
				c: &Permission{
					Verbs:         util.NewStringSetFromSlice([]string{"create"}),
					ResourceNames: util.NewStringSetFromSlice([]string{"test", "test2"}),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Permission{
				Verbs:         tt.fields.Verbs,
				ResourceNames: tt.fields.ResourceNames,
			}
			if got := p.Contains(tt.args.c); got != tt.want {
				t.Errorf("Permission.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
