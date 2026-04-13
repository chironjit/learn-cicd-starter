package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		wantAPIKey  string
		wantErr     bool
		wantErrText string
	}{
		{
			name:       "returns api key from valid authorization header",
			headers:    http.Header{"Authorization": []string{"ApiKey super-secret-key"}},
			wantAPIKey: "super-secret-key",
		},
		{
			name:        "returns error when authorization header missing",
			headers:     http.Header{},
			wantErr:     true,
			wantErrText: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name:        "returns error for malformed authorization header",
			headers:     http.Header{"Authorization": []string{"Bearer super-secret-key"}},
			wantErr:     true,
			wantErrText: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAPIKey, err := GetAPIKey(tt.headers)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if err.Error() != tt.wantErrText {
					t.Fatalf("expected error %q, got %q", tt.wantErrText, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if gotAPIKey != tt.wantAPIKey {
				t.Fatalf("expected API key %q, got %q", tt.wantAPIKey, gotAPIKey)
			}
		})
	
}
