package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/cut4cut/toimi-test-work/internal/entity"
	"github.com/cut4cut/toimi-test-work/internal/usecase"
	"github.com/cut4cut/toimi-test-work/pkg/logger"
)

type advertRoutes struct {
	u usecase.AdvertUseCase
	l logger.Interface
}

func newAdvertRoutes(handler *gin.RouterGroup, u usecase.AdvertUseCase, l logger.Interface) {
	r := &advertRoutes{u, l}

	h := handler.Group("/advert")
	{
		h.POST("/", r.create)
		h.GET("/:id", r.getById)
		h.GET("/page/", r.getPage)
	}
}

type correctResponse struct {
	Data interface{} `json:"data"`
}

func (r *advertRoutes) create(c *gin.Context) {
	adv := entity.NewAdvert()
	err := c.BindJSON(adv)
	if err != nil {
		r.l.Error(err, "http - v1 - create")
		errorResponse(c, http.StatusBadRequest, "bind JSON error")
		return
	}

	id, err := r.u.Create(c.Request.Context(), adv)
	if err != nil {
		r.l.Error(err, "http - v1 - create")
		errorResponse(c, http.StatusInternalServerError, "create advert error")
		return
	}

	c.JSON(http.StatusOK, correctResponse{id})
}

func (r *advertRoutes) getById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		r.l.Error(err, "http - v1 - getById")
		errorResponse(c, http.StatusBadRequest, "incorrect account ID")
		return
	}

	adv, err := r.u.GetById(c.Request.Context(), id)
	if err != nil {
		r.l.Error(err, "http - v1 - getById")
		errorResponse(c, http.StatusInternalServerError, "can not find advert")
		return
	}

	c.JSON(http.StatusOK, correctResponse{adv})
}

func (r *advertRoutes) getPage(c *gin.Context) {
	pag := entity.NewPagination()
	if err := c.ShouldBindQuery(&pag); err != nil {
		r.l.Error(err, "http - v1 - getPage")
		errorResponse(c, http.StatusBadRequest, "param query error")
		return
	}

	advs, err := r.u.GetPage(c.Request.Context(), &pag)
	if err != nil {
		r.l.Error(err, "http - v1 - getPage")
		errorResponse(c, http.StatusInternalServerError, "some error in usecase")
		return
	}

	c.JSON(http.StatusOK, correctResponse{advs})
}
