package e2e

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/domains"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/ports"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/services"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos"
	BookRepo "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos/book"
	router "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers"
	App "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers/app"
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
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/book", nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
	suite.Equal("404 page not found", w.Body.String())
}

func (suite *TestSuite) TestUseWrongParamType() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/book/one", nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
	suite.JSONEq(`{
		"msg": "strconv.ParseUint: parsing \"one\": invalid syntax"
	  }`, w.Body.String())
}

func (suite *TestSuite) TestFindNotFondBookByID() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/book/2", nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.JSONEq(`{
		"msg": "Book Not Found"
	  }`, w.Body.String())
}

func (suite *TestSuite) TestFoundBookID() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/book/1", nil)
	suite.router.ServeHTTP(w, req)

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

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(string(want), w.Body.String())
}

/*TestRoutePingSuite is trigger run it test*/
func TestRouteBookSuite(t *testing.T) {
	err := godotenv.Load("../../.env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	suite.Run(t, new(TestSuite))
}
