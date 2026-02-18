package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h1 := http.Header{}
	h1.Set("Authorization", "ApiKey 1234567890")
	h2 := http.Header{}
	h2.Set("Authorization", "Bearer 1234567890")
	h3 := http.Header{}
	h3.Set("Authorization", "ApiKey1234567890")

	tests := []struct {
		name     string
		header   http.Header
		expected string
		wantErr  bool
	}{
		{
			name:     "valid header",
			header:   h1,
			expected: "1234567890",
			wantErr:  false,
		},
		{
			name:     "empty invalid header",
			header:   http.Header{},
			expected: "",
			wantErr:  true,
		},
		{
			name:     "malformed header with ApiKey",
			header:   h2,
			expected: "",
			wantErr:  true,
		},
		{
			name:     "malformed header with no spaces",
			header:   h3,
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GetAPIKey(tt.header)
			if actual != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, actual)
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
