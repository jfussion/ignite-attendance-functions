package usecase_test

import (
	"context"
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
		ID:     "IGNT-DEMO-0001",
		Name:   "Joe",
		School: "PUP",
		Course: "BSCpE",
	}

	t.Run("success", func(t *testing.T) {
		mockPeopleRepo.On("Get", mock.Anything, mock.AnythingOfType("string")).Return(tPeople, nil).Once()
		id := "IGNT-DEMO-0001"
		got, err := mockPeopleUcase.Get(context.TODO(), id)
		assert.Equal(t, tPeople, got)
		assert.NoError(t, err)
		mockPeopleRepo.AssertExpectations(t)
	})

}
