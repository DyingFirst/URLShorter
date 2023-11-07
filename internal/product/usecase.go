package product

type UseCase interface {
	NewShort(OriginalURL string) (ShortedURL string, err error)
	GetOriginalURL(ShortedURL string) (OriginalURL string, err error)
}
