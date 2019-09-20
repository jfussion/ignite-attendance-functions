package domain

import "context"

type Attendance struct {
	Date   string
	People People
}

type Count struct {
	Total, Members, VIPs int
}

type AttendanceUsecase interface {
	Add(ctx context.Context, attendance Attendance) (err error)
	IncrementCount(ctx context.Context, isMember bool) (err error)
	Update(ctx context.Context, id string, attendance Attendance) (err error)
}

type AttendanceRepository interface {
	Add(ctx context.Context, attendance Attendance) (err error)
	GetCount(ctx context.Context, id string) (count Count, err error)
	UpdateCount(ctx context.Context, id string, count Count) (err error)
	Update(ctx context.Context, id string, attendance Attendance) (err error)
}
