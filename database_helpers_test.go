package parser

import (
	"testing"
)

func Test_checkTag(t *testing.T) {
	type args struct {
		fieldTags []string
		tag       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Check Required Tag",
			args: args{
				fieldTags: []string{
					propertyRequired,
					propertyIndex,
					propertyUnique,
				},
				tag: propertyRequired,
			},
			want: true,
		},
		{
			name: "Check Index Tag",
			args: args{
				fieldTags: []string{
					propertyRequired,
					propertyIndex,
					propertyUnique,
				},
				tag: propertyIndex,
			},
			want: true,
		},
		{
			name: "Check Unique Tag",
			args: args{
				fieldTags: []string{
					propertyRequired,
					propertyIndex,
					propertyUnique,
				},
				tag: propertyUnique,
			},
			want: true,
		},
		{
			name: "Check Invalid Tag",
			args: args{
				fieldTags: []string{
					propertyRequired,
					propertyIndex,
					propertyUnique,
				},
				tag: "invalid",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkTag(tt.args.fieldTags, tt.args.tag); got != tt.want {
				t.Errorf("checkTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFieldname(t *testing.T) {
	type args struct {
		fieldTags []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Check Required Field",
			args: args{
				fieldTags: []string{
					propertyRequired,
					propertyIndex,
					propertyUnique,
				},
			},
			want: "",
		},
		{
			name: "Check Index Field",
			args: args{
				fieldTags: []string{
					propertyRequired,
					propertyIndex,
					propertyUnique,
				},
			},
			want: "",
		},
		{
			name: "Check Unique Field",
			args: args{
				fieldTags: []string{
					propertyRequired,
					propertyIndex,
					propertyUnique,
				},
			},
			want: "",
		},
		{
			name: "Check Valid Field",
			args: args{
				fieldTags: []string{
					propertyRequired,
					propertyIndex,
					propertyUnique,
					"name",
				},
			},
			want: "name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFieldname(tt.args.fieldTags); got != tt.want {
				t.Errorf("getFieldname() = %v, want %v", got, tt.want)
			}
		})
	}
}
