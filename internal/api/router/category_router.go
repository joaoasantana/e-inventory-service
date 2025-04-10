package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoasantana/e-inventory-service/internal/api/handler"
	"github.com/joaoasantana/e-inventory-service/internal/domain/usecase"
)

const (
	relativePATH = "categories"
)

func RegisterCategory(api *gin.RouterGroup, category *usecase.CategoryUseCase) {
	h := handler.NewCategoryHandler(category)

	categoriesAPI := api.Group(relativePATH)
	{
		categoriesAPI.POST("/", h.CreateCategory)
		categoriesAPI.GET("/", h.FetchAllCategories)
		categoriesAPI.GET("/:id", h.FetchCategoryByID)
	}
}
