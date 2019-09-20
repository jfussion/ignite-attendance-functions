package domain

import "context"

type People struct {
	ID, Name, School, Course string
	IsMember                 bool
}

type PeopleUsecase interface {
	Get(ctx context.Context, id string) (people People, err error)
}

type PeopleRepository interface {
	Get(ctx context.Context, id string) (people People, err error)
}
