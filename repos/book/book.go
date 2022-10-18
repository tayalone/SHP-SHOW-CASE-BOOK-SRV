package book

import (
	"errors"

	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/domains"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos"
)

/*Repository (Adaptor) is Definition of Value */
type Repository struct {
	db *repos.RDB
}

var bookRepo = Repository{}

/*New do Create Rdb Connection*/
func New(db *repos.RDB) *Repository {
	bookRepo.db = db
	return &bookRepo
}

/*GetByPk Find Book in Database by Pk */
func (r *Repository) GetByPk(id uint) (domains.Book, error) {
	var b domains.Book

	result := r.db.First(&b, id)

	if result.RowsAffected != 1 {
		return domains.Book{}, errors.New("Barcode Condition Not Found")
	}
	return b, nil
}
