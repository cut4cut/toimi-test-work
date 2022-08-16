package v1

import (
	"github.com/cut4cut/toimi-test-work/internal/usecase"
	"github.com/cut4cut/toimi-test-work/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, l logger.Interface, u usecase.AdvertUseCase) {

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/v1")
	{
		newAdvertRoutes(h, u, l)
	}
}
