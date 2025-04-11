package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joaoasantana/e-inventory-service/internal/api/handler"
	"github.com/joaoasantana/e-inventory-service/internal/domain/usecase"
	"github.com/joaoasantana/e-inventory-service/internal/infra/repository"
	"go.uber.org/zap"
)

const categoryRoute = "/categories"

func InitCategoryRoute(db *sqlx.DB, logger *zap.Logger, api *gin.RouterGroup) {
	categoryUseCase := &usecase.CategoryUseCase{
		Logger:     logger,
		Repository: repository.NewCategoryRepository(db),
	}

	categoryHandler := handler.CategoryHandler{
		Logger:  logger,
		UseCase: categoryUseCase,
	}

	categoriesAPI := api.Group(categoryRoute)
	{
		categoriesAPI.POST("/", categoryHandler.CreateCategory)

		categoriesAPI.GET("/", categoryHandler.FetchAllCategories)
		categoriesAPI.GET("/:id", categoryHandler.FetchCategoryByID)
	}
}
