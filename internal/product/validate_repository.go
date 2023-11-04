package product

type ValidateRepo interface {
	ValidateURL(URL string) bool
}
