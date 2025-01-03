package parser

import (
	"reflect"
	"testing"
)

func TestEncodeModel(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    M
		wantErr bool
	}{
		{
			name: "Test Empty Struct Tags",
			args: args{
				obj: struct {
					Name string
				}{
					Name: "Jon",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test Invalid Struct Tags",
			args: args{
				obj: struct {
					Name string `tag:"name"`
				}{
					Name: "Jon",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test Valid Struct Tags",
			args: args{
				obj: struct {
					Name string `db:"name,required,unique,index"`
				}{
					Name: "Jon",
				},
			},
			want: M{
				{
					Key:      "name",
					Value:    "Jon",
					Required: true,
					Unique:   true,
					Index:    true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeModel(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
