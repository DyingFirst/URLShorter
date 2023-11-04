package validate_repository

import "testing"

func TestValidateURL_ValidURL(t *testing.T) {
	repo := NewValidateRepo()

	validURLs := []string{
		"https://www.example.com",
		"http://example.com",
		"http://sub.example.com/page",
	}

	for _, url := range validURLs {
		if !repo.ValidateURL(url) {
			t.Errorf("Expected URL to be valid: %s", url)
		}
	}
}

func TestValidateURL_InvalidURL(t *testing.T) {
	repo := NewValidateRepo()

	invalidURLs := []string{
		"ftp://example.com",
		"htp://example.com",
		"example.com",
		"https://.example.com",
	}

	for _, url := range invalidURLs {
		if repo.ValidateURL(url) {
			t.Errorf("Expected URL to be invalid: %s", url)
		}
	}
}
