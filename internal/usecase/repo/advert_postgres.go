package repo

import (
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/cut4cut/toimi-test-work/internal/entity"
	"github.com/cut4cut/toimi-test-work/pkg/postgres"
	"github.com/georgysavva/scany/pgxscan"
)

type AdvertRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *AdvertRepo {
	return &AdvertRepo{pg}
}

func newArrayExpr(items []string) string {
	if len(items) == 0 {
		return "DEFAULT"
	}

	return fmt.Sprintf("ARRAY['%s']", strings.Join(items, "', '"))
}

func newSortExpr(pag *entity.Pagination) (string, string) {
	sortPrice, sortCreated := "price ASC", "created_dt ASC"

	if (*pag).DescPrice {
		sortPrice = "price DESC"
	}
	if (*pag).DescCreated {
		sortCreated = "created_dt DESC"
	}

	return sortPrice, sortCreated
}

func (r *AdvertRepo) Create(ctx context.Context, adv *entity.Advert) (id int64, err error) {
	arrayExpr := newArrayExpr((*adv).Urls)
	sql, _, err := r.Builder.
		Insert("advert").
		Columns("id, created_dt, name, description, price, urls").
		Values(
			sq.Expr("DEFAULT"),
			sq.Expr("DEFAULT"),
			(*adv).Name,
			(*adv).Description,
			(*adv).Price,
			sq.Expr(arrayExpr),
		).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return id, fmt.Errorf("AdvertRepo - Create - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql,
		(*adv).Name,
		(*adv).Description,
		(*adv).Price).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("AdvertRepo - Create - tx.QueryRow: %w", err)
	}

	return
}

func (r *AdvertRepo) GetById(ctx context.Context, id int64) (adv entity.Advert, err error) {
	sql, _, err := r.Builder.
		Select("name, price, urls[:1]").
		From("advert").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return adv, fmt.Errorf("AdvertRepo - GetByID - r.Builder: %w", err)
	}

	dst := make([]entity.Advert, 0)
	if err := pgxscan.Select(
		ctx, r.Pool, &dst, sql, id,
	); err != nil {
		return adv, fmt.Errorf("AdvertRepo - GetByID - pgxscan.Select: %w", err)
	}

	return dst[0], nil
}

func (r *AdvertRepo) GetPage(ctx context.Context, pag *entity.Pagination) (advs []entity.Advert, err error) {
	offset := (*pag).Page * (*pag).ItemsPerPage
	sortExprPrice, sortExprCreated := newSortExpr(pag)
	sql, _, err := r.Builder.
		Select("name, price, urls[:1]").
		From("advert").
		OrderBy(sortExprPrice).
		OrderBy(sortExprCreated).
		Limit((*pag).ItemsPerPage).
		Offset(offset).
		ToSql()
	if err != nil {
		return advs, fmt.Errorf("AdvertRepo - GetHistory - r.Builder: %w", err)
	}

	advs = make([]entity.Advert, int((*pag).ItemsPerPage))
	if err := pgxscan.Select(
		ctx, r.Pool, &advs, sql,
	); err != nil {
		return nil, fmt.Errorf("AdvertRepo - GetHistory - pgxscan.Select: %w", err)
	}

	return
}
