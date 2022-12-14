package ports

import "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/domains"

/*BookSrv is Defining Behavior of "Book Service" */
type BookSrv interface {
	GetByID(id uint) (domains.Book, error)
}
