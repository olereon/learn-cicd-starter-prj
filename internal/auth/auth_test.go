package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError string
		wantErr       bool
	}{
		{
			name:          "valid api key",
			headers:       http.Header{"Authorization": []string{"ApiKey test-key-123"}},
			expectedKey:   "test-key-123",
			wantErr:       false,
		},
		{
			name:          "missing authorization header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: "no authorization header included",
			wantErr:       true,
		},
		{
			name:          "empty authorization header",
			headers:       http.Header{"Authorization": []string{""}},
			expectedKey:   "",
			expectedError: "no authorization header included",
			wantErr:       true,
		},
		{
			name:          "malformed header - wrong prefix",
			headers:       http.Header{"Authorization": []string{"Bearer test-key-123"}},
			expectedKey:   "",
			expectedError: "malformed authorization header",
			wantErr:       true,
		},
		{
			name:          "malformed header - no space",
			headers:       http.Header{"Authorization": []string{"ApiKeytest-key-123"}},
			expectedKey:   "",
			expectedError: "malformed authorization header",
			wantErr:       true,
		},
		{
			name:          "malformed header - only prefix",
			headers:       http.Header{"Authorization": []string{"ApiKey"}},
			expectedKey:   "",
			expectedError: "malformed authorization header",
			wantErr:       true,
		},
		{
			name:          "valid api key with extra spaces",
			headers:       http.Header{"Authorization": []string{"ApiKey test-key-with-spaces"}},
			expectedKey:   "test-key-with-spaces",
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)
			
			if key != tt.expectedKey {
				t.Errorf("expected key %q, got %q", tt.expectedKey, key)
			}
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error containing %q, got nil", tt.expectedError)
				} else if err.Error() != tt.expectedError {
					t.Errorf("expected error %q, got %q", tt.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			}
		})
	}
}