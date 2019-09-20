package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jfussion/ignite-attendance-cloud-functions/domain"
	"github.com/jfussion/ignite-attendance-cloud-functions/domain/mocks"
	"github.com/jfussion/ignite-attendance-cloud-functions/people/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPeople(t *testing.T) {
	mockPeopleRepo := new(mocks.PeopleRepository)
	mockPeopleUcase := usecase.NewUsecase(mockPeopleRepo)

	tPeople := domain.People{
		ID:       "IGNT-DEMO-0001",
		Name:     "Joe",
		School:   "PUP",
		Course:   "BSCpE",
		IsMember: true,
	}

	id := "IGNT-DEMO-0001"
	t.Run("success", func(t *testing.T) {
		mockPeopleRepo.On("Get", mock.Anything, id).Return(tPeople, nil).Once()
		got, err := mockPeopleUcase.Get(context.TODO(), id)
		assert.Equal(t, tPeople, got)
		assert.NoError(t, err)
		mockPeopleRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockPeopleRepo.On("Get", mock.Anything, mock.AnythingOfType("string")).Return(domain.People{}, errors.New("error: Something went wrong")).Once()
		got, err := mockPeopleUcase.Get(context.TODO(), id)

		assert.Empty(t, got)
		assert.Error(t, err)
		mockPeopleRepo.AssertExpectations(t)
	})
}
