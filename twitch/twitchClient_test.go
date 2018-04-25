package twitch

import (
	"io"
	"testing"
)

func TestRequest(t *testing.T) {
	type args struct {
		method  string
		url     string
		body    io.Reader
		auth    bool
		context bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get top games",
			args: args{
				method:  "GET",
				url:     "https://api.twitch.tv/kraken/games/top/",
				body:    nil,
				auth:    false,
				context: false,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Request(tt.args.method, tt.args.url, tt.args.body, tt.args.auth, tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) <= tt.want {
				t.Errorf("Request() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
