package app

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/cut4cut/toimi-test-work/config"
	v1 "github.com/cut4cut/toimi-test-work/internal/controller/http/v1"
	"github.com/cut4cut/toimi-test-work/internal/usecase"
	"github.com/cut4cut/toimi-test-work/internal/usecase/repo"
	"github.com/cut4cut/toimi-test-work/pkg/logger"
	"github.com/cut4cut/toimi-test-work/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	r := repo.New(pg)
	accountUseCase := usecase.New(r)

	handler := gin.Default()
	v1.NewRouter(handler, l, *accountUseCase)

	handler.Run()
}
