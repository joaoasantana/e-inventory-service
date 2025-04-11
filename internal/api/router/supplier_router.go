package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joaoasantana/e-inventory-service/internal/api/handler"
	"github.com/joaoasantana/e-inventory-service/internal/domain/usecase"
	"github.com/joaoasantana/e-inventory-service/internal/infra/repository"
	"go.uber.org/zap"
)

const supplierRoute = "/suppliers"

func InitSupplierRoute(db *sqlx.DB, logger *zap.Logger, api *gin.RouterGroup) {
	supplierUseCase := &usecase.SupplierUseCase{
		Logger:     logger,
		Repository: repository.NewSupplierRepository(db),
	}

	supplierHandler := handler.SupplierHandler{
		Logger:  logger,
		UseCase: supplierUseCase,
	}

	routerGroup := api.Group(supplierRoute)
	{
		routerGroup.POST("/", supplierHandler.CreateSupplier)

		routerGroup.GET("/", supplierHandler.FetchAllSuppliers)
		routerGroup.GET("/:id", supplierHandler.FetchSupplierByID)
	}
}
