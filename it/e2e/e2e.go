package e2e

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/ports"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos"
)

/*E2ETestSuite is a test suit for Repo*/
type E2ETestSuite struct {
	suite.Suite
	db   *repos.RDB
	repo ports.BookRpstr
}

/*TestRoutePingSuite is trigger run it test*/
func TestRoutePingSuite(t *testing.T) {
	err := godotenv.Load("../../.env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	suite.Run(t, new(E2ETestSuite))
}
