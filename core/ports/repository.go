package ports

import "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/domains"

/*BookRpstr is Define Behavior of "Book Repository" */
type BookRpstr interface {
	GetByPk(uint) (domains.Book, error)
}
