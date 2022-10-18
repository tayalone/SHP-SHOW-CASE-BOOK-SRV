package services

import (
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/domains"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/ports"
)

/*BookService is a Instance Of Book  */
type BookService struct {
	repo ports.BookRpstr
}

var srv = new(BookService)

/*
New is Return BookService Singleton Instance
https://refactoring.guru/design-patterns/singleton
*/
func New(r ports.BookRpstr) ports.BookSrv {
	srv.repo = r
	return srv
}

/*GetByID is Get Book from Repository By ID */
func (s *BookService) GetByID(id uint) (domains.Book, error) {
	b, err := s.repo.GetByPk(id)
	if err != nil {
		return domains.Book{}, err
	}
	return b, nil
}
