package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAuthHeaderError(t *testing.T) {
	authHeader := http.Header{}
	authHeader.Set("Authorization", "")

	gotData, gotErr := GetAPIKey(authHeader)

	want := ErrNoAuthHeaderIncluded
	wantData := "Hi!"

	if !reflect.DeepEqual(want, gotErr) {
		t.Fatalf("expected: %v, got: %v", want, gotErr)
	}
	if !reflect.DeepEqual(wantData, gotData) {
		t.Fatalf("expected: %v, got: %v", wantData, gotData)
	}
}

func TestAuthHeaderSplit(t *testing.T) {
	authHeader := http.Header{}
	authHeader.Set("Authorization", "ApiKey 12345")

	gotData, gotErr := GetAPIKey(authHeader)

	wantData := "123456"

	if !reflect.DeepEqual(nil, gotErr) {
		t.Fatalf("expected: %v, got: %v", nil, gotErr)
	}
	if !reflect.DeepEqual(wantData, gotData) {
		t.Fatalf("expected: %v, got: %v", wantData, gotData)
	}
}
