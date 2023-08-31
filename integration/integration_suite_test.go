package integration_test

import (
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http/httptest"
	"testing"
	"todo/db"
	"todo/server"
)

var database *sqlx.DB
var tx *sqlx.Tx

var testServer *httptest.Server

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = BeforeSuite(func() {
	database = db.Setup("postgres://hellofreshdev:hellofreshdev@localhost:5432/todo_test?sslmode=disable")
})

var _ = BeforeEach(func() {
	tx = database.MustBegin()
	router := server.Setup(tx)
	testServer = httptest.NewServer(router)
})

var _ = AfterEach(func() {
	Expect(tx.Rollback()).To(Succeed())
	testServer.Close()
})

var _ = AfterSuite(func() {
	Expect(database.Close()).To(Succeed())
})
