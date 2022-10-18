package repos

import (
	"fmt"
	"log"
	"os"

	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/domains"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*RDB is Definition of Value */
type RDB struct {
	*gorm.DB
	errMsg string
}

var rdb = RDB{
	DB:     nil,
	errMsg: "",
}

func migrate(db *gorm.DB) {
	// db.Set("gorm:table_options", "ENGINE=InnoDB")

	// /  about 'barcode_condition'
	if (db.Migrator().HasTable(&domains.Book{})) {
		log.Println("Table Existing, Drop IT")

		db.Migrator().DropTable(&domains.Book{})
	}
	db.AutoMigrate(&domains.Book{})

	var desc = "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s"

	// / Add Initail Data
	initialBook := []domains.Book{
		{
			Title:  "Lorem",
			Desc:   &desc,
			Author: "Dante Allergie",
		},
	}
	db.Create(initialBook)

	log.Println("Create 'books'")
}

/*New do Create Rdb Connection*/
func New() *RDB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("RDM_HOST"),
		os.Getenv("RDM_USER"),
		os.Getenv("RDM_PASSWORD"),
		os.Getenv("RDM_DB"),
		os.Getenv("RDM_PORT"),
		os.Getenv("TIME_ZONE"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		log.Println("FAIL: Connect RDB Error", err.Error())

		rdb.errMsg = err.Error()
		return &rdb

	}
	log.Println("Connect RDB Success!!!")
	rdb.DB = db
	rdb.errMsg = ""

	migrate(db)

	return &rdb
}
