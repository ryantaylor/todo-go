package user_test

import (
	"github.com/jmoiron/sqlx"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var db *sqlx.DB
var tx *sqlx.Tx

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}

var _ = BeforeSuite(func() {
	db = sqlx.MustConnect("pgx", "postgres://hellofreshdev:hellofreshdev@localhost:5432/todo_test?sslmode=disable")
})

var _ = BeforeEach(func() {
	tx = db.MustBegin()
})

var _ = AfterEach(func() {
	Expect(tx.Rollback()).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(db.Close()).To(Succeed())
})
