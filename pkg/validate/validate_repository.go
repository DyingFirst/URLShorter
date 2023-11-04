package validate

import (
	"regexp"
)

func ValidateURL(URL string) bool {
	regexPattern := `^(https?|http):\/\/(-\.)?([^\s/?\.#-]+\.?)+(\/[^\s]*)?$`
	regex := regexp.MustCompile(regexPattern)
	return regex.MatchString(URL)
}
