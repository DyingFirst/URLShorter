package validate_repository

import (
	"URLShorter/internal/product"
	"regexp"
)

type ValidateRepo struct {
}

func NewValidateRepo() product.ValidateRepo {
	return &ValidateRepo{}
}

func (val *ValidateRepo) ValidateURL(URL string) bool {
	regexPattern := `^(https?|http):\/\/(-\.)?([^\s/?\.#-]+\.?)+(\/[^\s]*)?$`
	regex := regexp.MustCompile(regexPattern)
	return regex.MatchString(URL)
}
