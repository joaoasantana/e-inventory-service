package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoasantana/e-inventory-service/internal/api/handler"
	"github.com/joaoasantana/e-inventory-service/internal/domain/usecase"
)

const (
	relativePath = "/products"
)

func RegisterProductRouter(api *gin.RouterGroup, product *usecase.ProductUseCase) {
	h := handler.NewProductHandler(product)

	productsAPI := api.Group(relativePath)
	{
		productsAPI.POST("/", h.CreateProduct)
		productsAPI.GET("/", h.FetchAllProducts)
		productsAPI.GET("/:id", h.FetchProductByID)
	}
}
