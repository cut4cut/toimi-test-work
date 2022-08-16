package usecase

import (
	"context"

	"github.com/cut4cut/toimi-test-work/internal/entity"
)

type AdvertRepo interface {
	Create(context.Context, *entity.Advert) (int64, error)
	GetById(context.Context, int64) (entity.Advert, error)
	GetPage(context.Context, *entity.Pagination) ([]entity.Advert, error)
}
