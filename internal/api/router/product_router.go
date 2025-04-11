package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joaoasantana/e-inventory-service/internal/api/handler"
	"github.com/joaoasantana/e-inventory-service/internal/domain/usecase"
	"github.com/joaoasantana/e-inventory-service/internal/infra/repository"
	"go.uber.org/zap"
)

const productRoute = "/products"

func InitProductRoute(db *sqlx.DB, logger *zap.Logger, api *gin.RouterGroup) {
	productUseCase := &usecase.ProductUseCase{
		Logger:          logger,
		CategoryUseCase: repository.NewCategoryRepository(db),
		ProductUseCase:  repository.NewProductRepository(db),
	}

	productHandler := handler.ProductHandler{
		Logger:  logger,
		UseCase: productUseCase,
	}

	routerGroup := api.Group(productRoute)
	{
		routerGroup.POST("/", productHandler.CreateProduct)

		routerGroup.GET("/", productHandler.FetchAllProducts)
		routerGroup.GET("/:id", productHandler.FetchProductByID)
	}
}
