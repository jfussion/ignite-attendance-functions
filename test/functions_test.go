package test

import (
	"context"
	"testing"

	f "github.com/jfussion/ignite-attendance-cloud-functions"
	"github.com/jfussion/ignite-attendance-cloud-functions/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPeople(t *testing.T) {
	// mockFirestoreEvent := f.FirestoreEvent{}
	mockClient := new(mocks.ClientRepo)
	mockGetterSetter := new(mocks.GetterSetter)
	mockDocSnapshot := new(mocks.DocSnapshot)

	tID := "IGNT-DEMO-0001"
	tData := map[string]interface{}{
		"name":     "Joseph",
		"school":   "PUP",
		"course":   "BSCpE",
		"isMember": true,
	}

	tPeople := f.People{
		ID:       tID,
		Name:     tData["name"].(string),
		School:   tData["school"].(string),
		Course:   tData["course"].(string),
		IsMember: tData["isMember"].(bool),
	}

	t.Run("success", func(t *testing.T) {
		mockClient.On("Doc", mock.AnythingOfType("string")).Return(mockGetterSetter).Once()
		mockGetterSetter.On("Get", mock.Anything).Return(mockDocSnapshot, nil).Once()
		// mockGetterSetter.On("Set", mock.Anything, mock.Anything).Return(nil, nil).Once()
		mockDocSnapshot.On("Data").Return(tData).Once()

		r, err := f.GetPeople(context.TODO(), mockClient, tID)

		assert.NoError(t, err)
		assert.Equal(t, tPeople, r)
	})
}

func TestPopulateAttendance(t *testing.T) {

}
