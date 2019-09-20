package usecase

import (
	"context"

	"github.com/jfussion/ignite-attendance-cloud-functions/domain"
)

type peopleUsecase struct {
	repo domain.PeopleRepository
}

func NewUsecase(peopleRepo domain.PeopleRepository) domain.PeopleUsecase {
	return &peopleUsecase{
		repo: peopleRepo,
	}
}

func (p *peopleUsecase) Get(ctx context.Context, id string) (people domain.People, err error) {
	return p.repo.Get(ctx, id)
}
