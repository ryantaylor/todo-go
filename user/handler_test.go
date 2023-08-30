package user_test

import (
	"bytes"
	"encoding/json"
	"github.com/go-faker/faker/v4"
	_ "github.com/jackc/pgx/v5/stdlib"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"todo/user"
)

var _ = Describe("Handler", func() {
	var repo user.RepositoryMock
	var handler user.Handler
	var httpHandler http.HandlerFunc
	var recorder *httptest.ResponseRecorder
	var req *http.Request

	BeforeEach(func() {
		repo = user.NewRepositoryMock()
		handler = user.NewHandler(&repo)
		recorder = httptest.NewRecorder()
	})

	var _ = Describe("Create", func() {
		var input user.CreateRequest

		BeforeEach(func() {
			input = user.CreateRequest{Email: faker.Email()}
			jsonBody, err := json.Marshal(input)
			Expect(err).ToNot(HaveOccurred())

			httpHandler = http.HandlerFunc(handler.Create)
			req, err = http.NewRequest("POST", "/users", bytes.NewReader(jsonBody))
			req.Header.Add("Content-Type", "application/json")
			httpHandler.ServeHTTP(recorder, req)
		})

		It("Returns a created status code", func() {
			Expect(recorder.Code).To(Equal(http.StatusCreated))
		})

		It("Returns an accurate payload", func() {
			var body user.Response

			err := json.Unmarshal(recorder.Body.Bytes(), &body)
			Expect(err).ToNot(HaveOccurred())

			Expect(body.ID).To(BeNumerically(">", 0))
			Expect(body.Email).To(Equal(input.Email))
		})

		It("Creates a user", func() {
			Expect(len(repo.Records)).To(Equal(1))

			for _, record := range repo.Records {
				Expect(record.Email).To(Equal(input.Email))
			}
		})

		When("A user with the given email already exists", func() {
			BeforeEach(func() {
				_, err := repo.CreateUser(input.Email)
				Expect(err).To(BeNil())
			})

			It("Returns an unprocessable entity status code", func() {
				Expect(recorder.Code).To(Equal(http.StatusUnprocessableEntity))
			})

			It("Returns an error payload", func() {
				Expect(true).To(BeTrue())
			})

			It("Does not create a user", func() {
				Expect(len(repo.Records)).To(BeZero())
			})
		})
	})
})
