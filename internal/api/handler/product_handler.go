package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/joaoasantana/e-inventory-service/internal/api/dto"
	"github.com/joaoasantana/e-inventory-service/internal/domain/model"
	"github.com/joaoasantana/e-inventory-service/internal/domain/usecase"
	"github.com/joaoasantana/e-inventory-service/pkg/utils"
	"go.uber.org/zap"
	"net/http"
)

type ProductHandler struct {
	Logger  *zap.Logger
	UseCase *usecase.ProductUseCase
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var requestBody dto.ProductRequest

	if err := ctx.ShouldBind(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Status: utils.StatusResponse{
				Code:    http.StatusBadRequest,
				Message: "invalid request body",
			},
			Error: err.Error(),
		})
		return
	}

	productModel := &model.Product{
		CategoryID:  requestBody.CategoryID,
		Name:        requestBody.Name,
		Image:       requestBody.Image,
		Price:       requestBody.Price,
		Description: requestBody.Description,
	}

	productID, err := h.UseCase.Create(productModel)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status: utils.StatusResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error while creating product",
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, utils.SuccessResponse{
		Status: utils.StatusResponse{
			Code:    http.StatusCreated,
			Message: "Product created",
		},
		Data: productID,
	})
}

func (h *ProductHandler) FetchAllProducts(ctx *gin.Context) {
	products, err := h.UseCase.FetchAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status: utils.StatusResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error while fetching products",
			},
			Error: err.Error(),
		})
		return
	}

	var result []dto.ProductResponse

	for _, product := range products {
		result = append(result, dto.ProductResponse{
			UUID:        product.UUID,
			Category:    product.CategoryID,
			Name:        product.Name,
			Image:       product.Image,
			Price:       product.Price,
			Description: product.Description,
		})
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse{
		Status: utils.StatusResponse{
			Code:    http.StatusOK,
			Message: "Products fetched",
		},
		Data: result,
	})
}

func (h *ProductHandler) FetchProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Status: utils.StatusResponse{
				Code:    http.StatusBadRequest,
				Message: "invalid request body",
			},
			Error: errors.New("missing product id").Error(), // todo
		})
		return
	}

	product, err := h.UseCase.FetchByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status: utils.StatusResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error while fetching product",
			},
			Error: err.Error(),
		})
		return
	}

	categoryResponse := dto.CategoryResponse{
		UUID:        product.Category.UUID,
		Name:        product.Category.Name,
		Description: product.Category.Description,
	}

	result := dto.ProductDetailResponse{
		UUID:        product.UUID,
		Category:    categoryResponse,
		Name:        product.Name,
		Image:       product.Image,
		Price:       product.Price,
		Description: product.Description,
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse{
		Status: utils.StatusResponse{
			Code:    http.StatusOK,
			Message: "Product fetched",
		},
		Data: result,
	})
}
