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

type CategoryHandler struct {
	Logger  *zap.Logger
	UseCase *usecase.CategoryUseCase
}

func (h *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var requestBody dto.CategoryRequest

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

	modelCategory := &model.Category{
		Name:        requestBody.Name,
		Description: requestBody.Description,
	}

	categoryID, err := h.UseCase.Create(modelCategory)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status: utils.StatusResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error while creating category",
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, utils.SuccessResponse{
		Status: utils.StatusResponse{
			Code:    http.StatusCreated,
			Message: "Category created",
		},
		Data: categoryID,
	})
}

func (h *CategoryHandler) FetchAllCategories(ctx *gin.Context) {
	categories, err := h.UseCase.FetchAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status: utils.StatusResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error while fetching categories",
			},
			Error: err.Error(),
		})
		return
	}

	var result []dto.CategoryResponse

	for _, category := range categories {
		result = append(result, dto.CategoryResponse{
			UUID:        category.UUID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse{
		Status: utils.StatusResponse{
			Code:    http.StatusOK,
			Message: "Categories fetched",
		},
		Data: result,
	})
}

func (h *CategoryHandler) FetchCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Status: utils.StatusResponse{
				Code:    http.StatusBadRequest,
				Message: "invalid request body",
			},
			Error: errors.New("missing category id").Error(), // todo
		})
		return
	}

	category, err := h.UseCase.FetchByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status: utils.StatusResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error while fetching category",
			},
			Error: err.Error(),
		})
		return
	}

	result := dto.CategoryResponse{
		UUID:        category.UUID,
		Name:        category.Name,
		Description: category.Description,
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse{
		Status: utils.StatusResponse{
			Code:    http.StatusOK,
			Message: "Category fetched",
		},
		Data: result,
	})
}
