package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseBody(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr string
	}{
		{
			name: "GET",
			args: args{
				r: httptest.NewRequest(
					http.MethodGet,
					"http://example.com/account",
					nil,
				),
			},
			want:    nil,
			wantErr: "",
		},
		{
			name: "DELETE",
			args: args{
				r: httptest.NewRequest(
					http.MethodDelete,
					"http://example.com/account/517b2d0d-d589-4f04-ae43-695579d24764",
					nil,
				),
			},
			want:    nil,
			wantErr: "",
		},
		{
			name: "POST no body",
			args: args{
				r: httptest.NewRequest(
					http.MethodPost,
					"http://example.com/account",
					nil,
				),
			},
			want:    nil,
			wantErr: "unexpected end of JSON input",
		},
		{
			name: "success",
			args: args{
				r: httptest.NewRequest(
					http.MethodPost,
					"http://example.com/account",
					strings.NewReader(`
					{
						"account_id": "test",
						"mail": "test@example.com"
					}
				`),
				),
			},
			want: map[string]interface{}{
				"account_id": "test",
				"mail":       "test@example.com",
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBody(tt.args.r)
			if err != nil {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parseIdentifier(t *testing.T) {
	type args struct {
		r        *http.Request
		endpoint string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "not specified id",
			args: args{
				r: httptest.NewRequest(
					http.MethodGet,
					"http://example.com/account",
					nil,
				),
				endpoint: "/account",
			},
			want: "",
		},
		{
			name: "surrounded by slashes",
			args: args{
				r: httptest.NewRequest(
					http.MethodGet,
					"http://example.com/account/c63d08f4-4728-4524-8672-1346d8d418c9",
					nil,
				),
				endpoint: "/account/",
			},
			want: "c63d08f4-4728-4524-8672-1346d8d418c9",
		},
		{
			name: "slash after endpoint",
			args: args{
				r: httptest.NewRequest(
					http.MethodGet,
					"http://example.com/account/c63d08f4-4728-4524-8672-1346d8d418c9",
					nil,
				),
				endpoint: "account/",
			},
			want: "c63d08f4-4728-4524-8672-1346d8d418c9",
		},
		{
			name: "no slash endpoint",
			args: args{
				r: httptest.NewRequest(
					http.MethodGet,
					"http://example.com/account/c63d08f4-4728-4524-8672-1346d8d418c9",
					nil,
				),
				endpoint: "account",
			},
			want: "c63d08f4-4728-4524-8672-1346d8d418c9",
		},
		{
			name: "success",
			args: args{
				r: httptest.NewRequest(
					http.MethodGet,
					"http://example.com/account/c63d08f4-4728-4524-8672-1346d8d418c9",
					nil,
				),
				endpoint: "/account",
			},
			want: "c63d08f4-4728-4524-8672-1346d8d418c9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseIdentifier(tt.args.r, tt.args.endpoint)
			assert.Equal(t, tt.want, got)
		})
	}
}
