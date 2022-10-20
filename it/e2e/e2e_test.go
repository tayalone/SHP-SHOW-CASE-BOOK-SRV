package e2e

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/domains"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/ports"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/services"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos"
	BookRepo "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos/book"
	App "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers/app"
	router "github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/router"
)

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

	log.Println("Mock Create 'books'")
}

/*TestSuite is a test suit for Repo*/
type TestSuite struct {
	suite.Suite
	db     *repos.RDB
	repo   ports.BookRpstr
	router router.Route
}

var (
	loc, _  = time.LoadLocation("Asia/Bangkok")
	tmpTime = time.Now().In(loc)
)

var desc = "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s"

/*SetupSuite init setup for BookRepo*/
func (suite *TestSuite) SetupSuite() {
	// suite.Router = router.SetUpRouter()

	db := repos.New()
	// ------- Make Repository
	bookRepo := BookRepo.New(db)

	suite.db = db
	initData(db)
	suite.repo = bookRepo
	bookSrv := services.New(bookRepo)
	myRouter := App.New(bookSrv)

	suite.router = myRouter
}

func (suite *TestSuite) TestNotUseQParams() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/book", nil)

	suite.Equal(http.StatusNotFound, statusCode)
	suite.Equal("404 page not found", actual)
}

func (suite *TestSuite) TestUseWrongParamType() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/book/one", nil)

	suite.Equal(http.StatusBadRequest, statusCode)
	suite.JSONEq(`{
		"msg": "strconv.ParseUint: parsing \"one\": invalid syntax"
	  }`, actual)
}

func (suite *TestSuite) TestFindNotFondBookByID() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/book/2", nil)

	suite.Equal(http.StatusInternalServerError, statusCode)
	suite.JSONEq(`{
		"msg": "Book Not Found"
	  }`, actual)
}

func (suite *TestSuite) TestFoundBookID() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/book/1", nil)

	wb := domains.Book{
		ID:        1,
		Title:     "Lorem",
		Author:    "Dante Allergie",
		Desc:      &desc,
		CreatedAt: tmpTime,
		UpdatedAt: tmpTime,
	}

	wantMap := map[string]interface{}{
		"book":    wb,
		"message": "OK",
	}

	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

/*TestRoutePingSuite is trigger run it test*/
func TestRouteBookSuite(t *testing.T) {
	err := godotenv.Load("../../.env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	suite.Run(t, new(TestSuite))
}
