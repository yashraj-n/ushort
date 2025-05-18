package services_test

import (
	"context"
	"testing"

	"github.com/joho/godotenv"
	"github.com/yashraj-n/ushort/services"
)

func TestCreateRedirectURL(t *testing.T) {
	loadEnv()
	tests := []struct {
		name   string
		url    string
		ipAddr string
	}{
		{
			name:   "Simple URL",
			url:    "http://google.com",
			ipAddr: "test.internal",
		},
		{
			name:   "Long URL",
			url:    "https://example.com/very/long/path/with/query?param=value&another=value",
			ipAddr: "test.internal",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := services.CreateRedirectURL(context.Background(), tt.url, tt.ipAddr)
			if err != nil {
				t.Errorf("CreateRedirectURL() error = %v", err)
				return
			}
			if hash == "" || len(hash) < 3 || len(hash) > 5 {
				t.Errorf("CreateRedirectURL() returned invalid hash: %v", hash)
			}
		})
	}
}

func TestGetRedirectURL(t *testing.T) {
	loadEnv()
	ctx := context.Background()

	// First create a URL to test with
	originalURL := "http://example.com"
	hash, err := services.CreateRedirectURL(ctx, originalURL, "test.internal")
	if err != nil {
		t.Fatalf("Failed to create test URL: %v", err)
	}

	// Test retrieving the URL
	got, err := services.GetRedirectURL(ctx, hash)
	if err != nil {
		t.Errorf("GetRedirectURL() error = %v", err)
		return
	}
	if got != originalURL {
		t.Errorf("GetRedirectURL() = %v, want %v", got, originalURL)
	}

	// Test cache hit
	got, err = services.GetRedirectURL(ctx, hash)
	if err != nil {
		t.Errorf("GetRedirectURL() cache hit error = %v", err)
		return
	}
	if got != originalURL {
		t.Errorf("GetRedirectURL() cache hit = %v, want %v", got, originalURL)
	}
}

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
}
