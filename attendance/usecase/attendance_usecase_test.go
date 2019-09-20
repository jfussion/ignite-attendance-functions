package usecase_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jfussion/ignite-attendance-cloud-functions/attendance/usecase"
	"github.com/jfussion/ignite-attendance-cloud-functions/domain"
	"github.com/jfussion/ignite-attendance-cloud-functions/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddAttendance(t *testing.T) {
	mockAttendanceRepo := new(mocks.AttendanceRepository)
	usecase := usecase.NewAttendanceUsecase(mockAttendanceRepo)
	tPeople := domain.People{
		ID:     "IGNT-DEMO-0001",
		Name:   "Joe",
		School: "PUP",
		Course: "BSCpE",
	}

	tAttendance := domain.Attendance{
		Date:   "September 19, 2019",
		People: tPeople,
	}

	t.Run("success", func(t *testing.T) {
		mockAttendanceRepo.On("Add", mock.Anything, tAttendance).Return(nil).Once()

		got := usecase.Add(context.TODO(), tAttendance)
		assert.NoError(t, got)
		mockAttendanceRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockAttendanceRepo.On("Add", mock.Anything, tAttendance).Return(errors.New("error: something went wrong")).Once()

		got := usecase.Add(context.TODO(), tAttendance)
		assert.Error(t, got)
		mockAttendanceRepo.AssertExpectations(t)
	})

}

func TestUpdateAttendance(t *testing.T) {
	mockAttendanceRepo := new(mocks.AttendanceRepository)
	attendanceUcase := usecase.NewAttendanceUsecase(mockAttendanceRepo)
	tPeople := domain.People{
		ID:     "IGNT-DEMO-0001",
		Name:   "Joe",
		School: "PUP",
		Course: "BSCpE",
	}

	tAttendance := domain.Attendance{
		Date:   "September 19, 2019",
		People: tPeople,
	}

	t.Run("success", func(t *testing.T) {
		mockAttendanceRepo.On("Update", mock.Anything, mock.AnythingOfType("string"), tAttendance).Return(nil).Once()

		err := attendanceUcase.Update(context.TODO(), "id", tAttendance)

		assert.NoError(t, err)
		mockAttendanceRepo.AssertExpectations(t)
	})
}

func TestIncrementCount(t *testing.T) {
	mockAttendanceRepo := new(mocks.AttendanceRepository)
	attendanceUcase := usecase.NewAttendanceUsecase(mockAttendanceRepo)
	tCount := domain.Count{
		Total:   2,
		Members: 1,
		VIPs:    1,
	}

	t.Run("success", func(t *testing.T) {
		mockAttendanceRepo.On("GetCount", mock.Anything, mock.AnythingOfType("string")).Return(tCount, nil).Once()
		mockAttendanceRepo.On("UpdateCount", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(nil).Once()

		got := attendanceUcase.IncrementCount(context.TODO(), true)
		assert.NoError(t, got)
		mockAttendanceRepo.AssertExpectations(t)
	})

	t.Run("success-increment members count", func(t *testing.T) {
		mockMembersIncremented := domain.Count{
			Total:   3,
			Members: 2,
			VIPs:    1,
		}

		mockAttendanceRepo.On("GetCount", mock.Anything, mock.AnythingOfType("string")).Return(tCount, nil).Once()
		mockAttendanceRepo.On("UpdateCount", mock.Anything, mock.AnythingOfType("string"), mockMembersIncremented).Return(nil).Once()

		got := attendanceUcase.IncrementCount(context.TODO(), true)
		assert.NoError(t, got)
		mockAttendanceRepo.AssertExpectations(t)
	})

	t.Run("success-increment VIPs count", func(t *testing.T) {
		mockVIPIncremented := domain.Count{
			Total:   3,
			Members: 1,
			VIPs:    2,
		}

		mockAttendanceRepo.On("GetCount", mock.Anything, mock.AnythingOfType("string")).Return(tCount, nil).Once()
		mockAttendanceRepo.On("UpdateCount", mock.Anything, mock.AnythingOfType("string"), mockVIPIncremented).Return(nil).Once()

		got := attendanceUcase.IncrementCount(context.TODO(), false)
		assert.NoError(t, got)
		mockAttendanceRepo.AssertExpectations(t)
	})

	t.Run("doc-ID", func(t *testing.T) {
		tDate := time.Now()
		id := fmt.Sprintf("count-%s", tDate.Format("2006-01-02"))

		mockAttendanceRepo.On("GetCount", mock.Anything, id).Return(tCount, nil).Once()
		mockAttendanceRepo.On("UpdateCount", mock.Anything, id, mock.Anything).Return(nil).Once()

		_ = attendanceUcase.IncrementCount(context.TODO(), true)
		mockAttendanceRepo.AssertExpectations(t)
	})

	t.Run("count doc does not exist", func(t *testing.T) {
		mockCount := domain.Count{
			Total:   1,
			Members: 1,
			VIPs:    0,
		}

		mockAttendanceRepo.On("GetCount", mock.Anything, mock.Anything).Return(domain.Count{}, nil).Once()
		mockAttendanceRepo.On("UpdateCount", mock.Anything, mock.Anything, mockCount).Return(nil).Once()
		err := attendanceUcase.IncrementCount(context.TODO(), true)
		assert.NoError(t, err)
	})
}
