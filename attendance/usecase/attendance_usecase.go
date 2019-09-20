package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/jfussion/ignite-attendance-cloud-functions/domain"
)

type attendaceUsecase struct {
	repo domain.AttendanceRepository
}

func NewAttendanceUsecase(attendanceRepo domain.AttendanceRepository) domain.AttendanceUsecase {
	return &attendaceUsecase{
		repo: attendanceRepo,
	}
}

func (a *attendaceUsecase) Add(ctx context.Context, attendance domain.Attendance) (err error) {
	return a.repo.Add(ctx, attendance)
}

func (a *attendaceUsecase) IncrementCount(ctx context.Context, isMember bool) (err error) {
	date := time.Now().Format("2006-01-02")
	id := fmt.Sprintf("count-%s", date)
	count, err := a.repo.GetCount(ctx, id)

	count.Total++
	if isMember {
		count.Members++
	} else {
		count.VIPs++
	}

	return a.repo.UpdateCount(ctx, id, count)
}

func (a *attendaceUsecase) Update(ctx context.Context, id string, attendance domain.Attendance) (err error) {
	return a.repo.Update(ctx, id, attendance)
}
