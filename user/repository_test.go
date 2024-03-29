package user_test

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/go-faker/faker/v4"
	_ "github.com/jackc/pgx/v5/stdlib"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	database "todo/db"
	"todo/models"
	"todo/user"
)

var _ = Describe("Repository", func() {
	var repo user.Repository
	var email string
	var err error

	BeforeEach(func() {
		repo = user.NewRepository(tx)
		email = faker.Email()
	})

	var _ = Describe("CreateUser", func() {
		It("Creates a user", func() {
			user, err := repo.CreateUser(email)
			Expect(err).To(BeNil())

			persistedUser, err := repo.FindUserByID(user.ID)
			Expect(err).To(BeNil())
			Expect(persistedUser.Email).To(Equal(email))
		})

		When("A user with the given email already exists", func() {
			BeforeEach(func() {
				_, err := repo.CreateUser(email)
				Expect(err).To(BeNil())
			})

			It("Returns an error", func() {
				_, err := repo.CreateUser(email)
				Expect(err).ToNot(BeNil())
			})

			It("Does not return a user", func() {
				user, _ := repo.CreateUser(email)
				Expect(user).To(BeNil())
			})

			It("Does not create a record", func() {
				var count int
				err = database.CountRecords(db, &count, database.TableUsers, sq.Eq{})
				Expect(count).To(BeZero())
			})
		})
	})

	var _ = Describe("FindUserByID", func() {
		var user *models.User

		BeforeEach(func() {
			user, err = repo.CreateUser(email)
			Expect(err).To(BeNil())
		})

		It("Returns the correct user", func() {
			queriedUser, err := repo.FindUserByID(user.ID)
			Expect(err).To(BeNil())
			Expect(queriedUser).To(Equal(user))
		})

		When("Given ID does not exist", func() {
			It("Returns an error", func() {
				_, err = repo.FindUserByID(-1)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
