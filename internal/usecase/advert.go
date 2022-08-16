package usecase

import (
	"context"
	"fmt"

	"github.com/cut4cut/toimi-test-work/internal/entity"
)

type AdvertUseCase struct {
	repo AdvertRepo
}

func New(r AdvertRepo) *AdvertUseCase {
	return &AdvertUseCase{
		repo: r,
	}
}

func (uc *AdvertUseCase) Create(ctx context.Context, adv *entity.Advert) (id int64, err error) {
	if err = (*adv).Check(); err != nil {
		return id, fmt.Errorf("AdvertUseCase - Create - validation: %w", err)
	}

	if id, err = uc.repo.Create(ctx, adv); err != nil {
		return id, fmt.Errorf("AdvertUseCase - Create - uc.repo.Create: %w", err)
	}

	return
}

func (uc *AdvertUseCase) GetById(ctx context.Context, id int64) (advResp entity.AdvertResponse, err error) {
	adv, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return advResp, fmt.Errorf("AdvertUseCase - GetById - uc.repo.GetById: %w", err)
	}

	advResp = entity.NewAdvertResponse(adv)

	return
}

func (uc *AdvertUseCase) GetPage(ctx context.Context, pag *entity.Pagination) (advResps []entity.AdvertResponse, err error) {
	advs, err := uc.repo.GetPage(ctx, pag)
	if err != nil {
		return advResps, fmt.Errorf("AdvertUseCase - GetPage - uc.repo.GetPage: %w", err)
	}

	length := len(advs)
	advResps = make([]entity.AdvertResponse, length)
	for i := 0; i < length; i++ {
		fmt.Println(i, advs)
		advResps[i] = entity.NewAdvertResponse(advs[i])
	}

	return
}
