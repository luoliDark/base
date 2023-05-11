package httputil

import (
	"net/url"
	"testing"
)

func TestPost(t *testing.T) {
	type args struct {
		apiURL string
		params url.Values
	}
	tests := []struct {
		name        string
		args        args
		wantResData string
		wantErr     bool
	}{
		//{
		//	name:"test1",
		//	args:args{
		//		apiURL:"http://127.0.0.1:9003/st/stinoroutbychk",
		//		params: map[string][]string{"pid":{"50201"},"primarykey":{"f3e8c06dc80845ad804aa4ef61f8284f"}},
		//	},
		//},
		{
			name: "test2",
			args: args{
				apiURL: "http://121.224.115.130:9011/base/loadedit",
				params: map[string][]string{"pid": {"50201"}, "primarykey": {"f3e8c06dc80845ad804aa4ef61f8284f"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResData, err := Post(tt.args.apiURL, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResData != tt.wantResData {
				t.Errorf("Post() gotResData = %v, want %v", gotResData, tt.wantResData)
			}
		})
	}
}
