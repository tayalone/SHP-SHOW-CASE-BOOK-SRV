package repos

import (
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/domains"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/ports"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos"
	BookRepo "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos/book"
)

/*BookRepositoryTestSuite is a test suit for Repo*/
type BookRepositoryTestSuite struct {
	suite.Suite
	db   *repos.RDB
	repo ports.BookRpstr
}

var loc, _ = time.LoadLocation("Asia/Bangkok")
var desc = "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s"
var tmpTime = time.Now().In(loc)

func initData(db *repos.RDB) {

	if (db.Migrator().HasTable(&domains.Book{})) {
		log.Println("Table Existing, Drop IT")

		db.Migrator().DropTable(&domains.Book{})
	}
	db.AutoMigrate(&domains.Book{})

	// / Add Initail Data
	initialBook := []domains.Book{
		{
			ID:        1,
			Title:     "Lorem",
			Desc:      &desc,
			Author:    "Dante Allergie",
			CreatedAt: tmpTime,
			UpdatedAt: tmpTime,
		},
	}
	db.Create(initialBook)

	log.Println("Create 'books'")
}

/*SetupSuite init setup for BookRepo*/
func (suite *BookRepositoryTestSuite) SetupSuite() {
	// suite.Router = router.SetUpRouter()
	db := repos.New()
	// ------- Make Repository
	bookRepo := BookRepo.New(db)

	suite.db = db
	suite.repo = bookRepo
	initData(db)
}

/*BeforeTest Do Setup Enviroment for each test case in suite */
func (suite *BookRepositoryTestSuite) BeforeTest(_, testName string) {
	// fmt.Println("BeforeTest name", testName)
}

/*TestFoundBookByID run Test Found Book By Pk */
func (suite *BookRepositoryTestSuite) TestFoundBookByID() {

	b, err := suite.repo.GetByPk(1)

	suite.EqualValues(b, domains.Book{
		ID:        1,
		Title:     "Lorem",
		Desc:      &desc,
		Author:    "Dante Allergie",
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	})
	suite.Equal(err, nil)
}

/*TestNotFoundBookByID run Test Not Found Book By Pk */
func (suite *BookRepositoryTestSuite) TestNotFoundBookByID() {
	want := domains.Book{}

	b, err := suite.repo.GetByPk(10)

	suite.EqualValues(b, want)
	suite.EqualError(err, "Book Not Found")

}

/*AfterTest Clear Enviroment in for each test case */
func (suite *BookRepositoryTestSuite) AfterTest(_, testName string) {
	// fmt.Println("AfterTest name", testName)
}

/*TestRoutePingSuite is trigger run it test*/
func TestRoutePingSuite(t *testing.T) {
	err := godotenv.Load("../../.env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	suite.Run(t, new(BookRepositoryTestSuite))
}
