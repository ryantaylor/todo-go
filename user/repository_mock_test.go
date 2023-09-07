package user

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"todo/db"
	"todo/models"
)

type RepositoryMock struct {
	Sequence       int
	RecordsByID    map[int]models.User
	RecordsByEmail map[string]models.User
}

func NewRepositoryMock() RepositoryMock {
	return RepositoryMock{
		Sequence:       1,
		RecordsByID:    make(map[int]models.User),
		RecordsByEmail: make(map[string]models.User),
	}
}

func (repo *RepositoryMock) CreateUser(email string) (*models.User, error) {
	var record models.User

	if _, exists := repo.RecordsByEmail[email]; exists {
		return nil, &db.DuplicateError{Message: fmt.Sprintf("A user with email %v already exists!", email)}
	}

	err := faker.FakeData(&record)
	if err != nil {
		return nil, err
	}

	record.ID = repo.Sequence
	record.Email = email

	repo.RecordsByID[record.ID] = record
	repo.RecordsByEmail[record.Email] = record
	repo.Sequence++

	return &record, nil
}

func (repo *RepositoryMock) FindUserByID(id int) (*models.User, error) {
	record, ok := repo.RecordsByID[id]
	if ok {
		return &record, nil
	}
	return nil, &db.NotFoundError{Message: fmt.Sprintf("No user with ID %v found!", id)}
}

func (repo *RepositoryMock) Clear() {
	repo.RecordsByID = make(map[int]models.User)
	repo.RecordsByEmail = make(map[string]models.User)
	repo.Sequence = 1
}
