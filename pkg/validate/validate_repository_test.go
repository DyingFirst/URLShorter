package validate

import "testing"

func TestValidateURL_ValidURL(t *testing.T) {

	validURLs := []string{
		"https://www.example.com",
		"http://example.com",
		"http://sub.example.com/page",
	}

	for _, url := range validURLs {
		if !ValidateURL(url) {
			t.Errorf("Expected URL to be valid: %s", url)
		}
	}
}

func TestValidateURL_InvalidURL(t *testing.T) {
	invalidURLs := []string{
		"ftp://example.com",
		"htp://example.com",
		"example.com",
		"https://.example.com",
	}

	for _, url := range invalidURLs {
		if ValidateURL(url) {
			t.Errorf("Expected URL to be invalid: %s", url)
		}
	}
}
