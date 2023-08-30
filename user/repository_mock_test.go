package user

import (
	"github.com/go-faker/faker/v4"
	"todo/models"
)

type RepositoryMock struct {
	Sequence int
	Records  map[int]models.User
}

func NewRepositoryMock() RepositoryMock {
	return RepositoryMock{
		Sequence: 1,
		Records:  make(map[int]models.User),
	}
}

func (repo *RepositoryMock) CreateUser(email string) (models.User, error) {
	var record models.User

	err := faker.FakeData(&record)
	if err != nil {
		return models.User{}, err
	}

	record.ID = repo.Sequence
	record.Email = email

	repo.Records[record.ID] = record
	repo.Sequence++

	return record, nil
}

func (repo *RepositoryMock) FindUserByID(id int) (models.User, error) {
	return repo.Records[id], nil
}

func (repo *RepositoryMock) Clear() {
	repo.Records = make(map[int]models.User)
	repo.Sequence = 1
}
