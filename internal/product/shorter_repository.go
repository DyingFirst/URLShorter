package product

type ShorterRepository interface {
	URLToID(url string) string
}
