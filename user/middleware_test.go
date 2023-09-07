package user_test

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strconv"
	"todo/models"
	"todo/user"
)

var _ = Describe("Middleware", func() {
	var repo user.RepositoryMock
	var middleware user.Middleware
	var err error

	BeforeEach(func() {
		repo = user.NewRepositoryMock()
		middleware = user.NewMiddleware(&repo)
	})

	Describe("Context", func() {
		var expectationHandler http.Handler
		var middlewareHandler http.Handler
		var req *http.Request
		var user *models.User

		JustBeforeEach(func() {
			req = httptest.NewRequest("GET", "http://testing", nil)
			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("userID", strconv.Itoa(user.ID))
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
		})

		When("User exists", func() {
			BeforeEach(func() {
				user, err = repo.CreateUser(faker.Email())
				Expect(err).ToNot(HaveOccurred())
			})

			It("Adds the user to the context", func() {
				expectationHandler = http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
					contextUser := req.Context().Value("user").(*models.User)
					Expect(contextUser).To(Equal(user))
				})

				middlewareHandler = middleware.Context(expectationHandler)
				middlewareHandler.ServeHTTP(httptest.NewRecorder(), req)
			})
		})

		When("User does not exist", func() {
			BeforeEach(func() {
				user = &models.User{ID: 1}
			})

			It("Does not add the user to the context", func() {
				expectationHandler = http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
					contextUser := req.Context().Value("user").(*models.User)
					Expect(contextUser).To(BeNil())
				})

				middlewareHandler = middleware.Context(expectationHandler)
				middlewareHandler.ServeHTTP(httptest.NewRecorder(), req)
			})
		})
	})
})
