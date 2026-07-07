package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  string
		want    string
		wantErr bool
	}{
		{"valid", "ApiKey 12345", "12345", false},
		{"empty", "", "", true},
		{"wrong_prefix", "Bearer 12345", "", true},
		{"missing_key", "ApiKey", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := http.Header{}
			if tc.header != "" {
				h.Set("Authorization", tc.header)
			}

			got, err := GetAPIKey(h)

			if (err != nil) != tc.wantErr {
				t.Fatalf("err: %v, wantErr: %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Fatalf("got: %q, want: %q", got, tc.want)
			}
		})
	}
}
