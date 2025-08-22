package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "ApiKey 123456test")
	result, error := GetAPIKey(headers)
	want := "123456test"

	if error != nil {
		t.Fatalf("expected: %v, got: %v", want, error)
	}

	t.Logf("expected: %v, got: %v", want, result)
}

// func TestGetAPIKeyBadKey(t *testing.T) {
// 	headers := make(http.Header)
// 	headers.Set("Authorization", "Bearer token123456")
// 	result, error := GetAPIKey(headers)
// 	want := "token123456"

// 	if error != nil {
// 		t.Fatalf("expected: %v, got: %v", want, error)
// 	}

// 	t.Logf("expected: %v, got: %v", want, result)
// }
