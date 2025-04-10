package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoasantana/e-inventory-service/internal/api/handler"
	"github.com/joaoasantana/e-inventory-service/internal/domain/usecase"
	"go.uber.org/zap"
)

const (
	relativePATH = "categories"
)

type CategoryRouter struct {
	logger   *zap.Logger
	category *usecase.CategoryUseCase
}

func NewCategoryRouter(logger *zap.Logger, category *usecase.CategoryUseCase) *CategoryRouter {
	return &CategoryRouter{logger, category}
}

func (r *CategoryRouter) Init(api *gin.RouterGroup) {
	h := handler.NewCategoryHandler(r.logger, r.category)

	categoriesAPI := api.Group(relativePATH)
	{
		categoriesAPI.POST("/", h.CreateCategory)
		categoriesAPI.GET("/", h.FetchAllCategories)
		categoriesAPI.GET("/:id", h.FetchCategoryByID)
	}
}
