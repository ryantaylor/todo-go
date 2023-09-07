package integration_test

import (
	"bytes"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"todo/db"
	"todo/models"
)

var _ = Describe("Users", func() {
	var err error
	var res *http.Response

	Describe("POST /users", func() {
		var email string

		BeforeEach(func() {
			email = faker.Email()
			body := []byte(fmt.Sprintf(`{ "email": "%v" }`, email))
			res, err = http.Post(fmt.Sprintf(`%s/users`, testServer.URL), "application/json", bytes.NewReader(body))
			Expect(err).ToNot(HaveOccurred())
		})

		It("Creates a single user", func() {
			var count int

			err = db.CountRecords(tx, &count, db.TableUsers, sq.Eq{})
			Expect(err).ToNot(HaveOccurred())
			Expect(count).To(Equal(1))
		})

		It("Creates the correct user", func() {
			var user models.User

			err = db.FindLastRecord(tx, &user, db.TableUsers)
			Expect(err).ToNot(HaveOccurred())
			Expect(user.Email).To(Equal(email))
		})

		It("Returns the correct status code", func() {
			Expect(res).To(HaveHTTPStatus(http.StatusCreated))
		})

		It("Returns the correct response body", func() {
			var user models.User

			err = db.FindLastRecord(tx, &user, db.TableUsers)
			Expect(err).ToNot(HaveOccurred())

			expected := fmt.Sprintf(`{
				"id": %v,
				"email": "%v"
			}`, user.ID, email)

			Expect(res).To(HaveHTTPBody(MatchJSON(expected)))
		})
	})
})
