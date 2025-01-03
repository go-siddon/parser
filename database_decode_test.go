package parser

import "testing"

func TestDecodeModel(t *testing.T) {
	type user struct {
		Jon string `db:"jon"`
	}
	var res user
	type args struct {
		obj  interface{}
		data M
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Non Pointer Struct",
			args: args{
				obj:  res,
				data: M{E{Key: "jon", Value: "jon"}}},
			wantErr: true,
		},
		{
			name: "Test Pointer Struct",
			args: args{
				obj:  &res,
				data: M{E{Key: "jon", Value: "jon"}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DecodeModel(tt.args.obj, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("DecodeModel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
