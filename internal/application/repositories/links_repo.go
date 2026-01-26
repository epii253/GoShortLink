package repositories

type ILinksRepo interface {
	CheckExsist(shrortLink string) bool

	TryAddItem(fullLink string, shrortLink string) bool
	DeleteItem(shrortLink string)

	GetByLink(shortLink string) (*string, bool)
}