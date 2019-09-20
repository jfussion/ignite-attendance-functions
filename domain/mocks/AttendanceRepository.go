// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/jfussion/ignite-attendance-cloud-functions/domain"
	mock "github.com/stretchr/testify/mock"
)

// AttendanceRepository is an autogenerated mock type for the AttendanceRepository type
type AttendanceRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, attendance
func (_m *AttendanceRepository) Add(ctx context.Context, attendance domain.Attendance) error {
	ret := _m.Called(ctx, attendance)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Attendance) error); ok {
		r0 = rf(ctx, attendance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCount provides a mock function with given fields: ctx, id
func (_m *AttendanceRepository) GetCount(ctx context.Context, id string) (domain.Count, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Count
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Count); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Count)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, attendance
func (_m *AttendanceRepository) Update(ctx context.Context, id string, attendance domain.Attendance) error {
	ret := _m.Called(ctx, id, attendance)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.Attendance) error); ok {
		r0 = rf(ctx, id, attendance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCount provides a mock function with given fields: ctx, id, count
func (_m *AttendanceRepository) UpdateCount(ctx context.Context, id string, count domain.Count) error {
	ret := _m.Called(ctx, id, count)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.Count) error); ok {
		r0 = rf(ctx, id, count)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}