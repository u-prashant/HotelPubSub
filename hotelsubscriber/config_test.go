package main

import "testing"

func Test_loadConfig(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "correct file path",
			args: args{
				file: "../hotelsubscriber/config.yaml",
			},
			wantErr: false,
		},
		{
			name: "wrong file path",
			args: args{
				file: "dummypath/config.yaml",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := loadConfig(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
