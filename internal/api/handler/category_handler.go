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
	logger          *zap.Logger
	categoryUseCase *usecase.CategoryUseCase
}

func NewCategoryHandler(logger *zap.Logger, categoryUseCase *usecase.CategoryUseCase) *CategoryHandler {
	mLogger := logger.With(
		zap.String("type", "handler"),
		zap.String("domain", "category"),
	)

	return &CategoryHandler{mLogger, categoryUseCase}
}

func (h *CategoryHandler) CreateCategory(ctx *gin.Context) {
	h.logger.Info("create", zap.String("status", "creating"), zap.String("method", "POST"))

	var requestBody dto.CategoryRequest

	if err := ctx.ShouldBind(&requestBody); err != nil {
		h.logger.Error("create", zap.String("status", "error"), zap.String("method", "POST"), zap.Error(err))

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

	categoryID, err := h.categoryUseCase.Create(modelCategory)
	if err != nil {
		h.logger.Error("create", zap.String("status", "error"), zap.String("method", "POST"), zap.Error(err))

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status: utils.StatusResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error while creating category",
			},
			Error: err.Error(),
		})
		return
	}

	h.logger.Info("create", zap.String("status", "created"), zap.String("method", "POST"))

	ctx.JSON(http.StatusCreated, utils.SuccessResponse{
		Status: utils.StatusResponse{
			Code:    http.StatusCreated,
			Message: "Category created",
		},
		Data: categoryID,
	})
}

func (h *CategoryHandler) FetchAllCategories(ctx *gin.Context) {
	h.logger.Info("fetchAll", zap.String("status", "fetching"), zap.String("method", "GET"))

	categories, err := h.categoryUseCase.FetchAll()
	if err != nil {
		h.logger.Error("fetchAll", zap.String("status", "error"), zap.String("method", "GET"), zap.Error(err))

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status: utils.StatusResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error while fetching categories",
			},
			Error: err.Error(),
		})
		return
	}

	h.logger.Info("fetchAll", zap.String("status", "mapping"), zap.String("method", "GET"), zap.Any("categories", categories))

	var result []dto.CategoryResponse

	for _, category := range categories {
		result = append(result, dto.CategoryResponse{
			UUID:        category.UUID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	h.logger.Info("fetchAll", zap.String("status", "fetched"), zap.String("method", "GET"), zap.Any("categories", result))

	ctx.JSON(http.StatusOK, utils.SuccessResponse{
		Status: utils.StatusResponse{
			Code:    http.StatusOK,
			Message: "Category fetched",
		},
		Data: result,
	})
}

func (h *CategoryHandler) FetchCategoryByID(ctx *gin.Context) {
	h.logger.Info("fetchByID", zap.String("status", "fetching"), zap.String("method", "GET"))

	id := ctx.Param("id")

	if id == "" {
		err := errors.New("missing category id")
		h.logger.Error("fetchByID", zap.String("status", "error"), zap.String("method", "GET"), zap.Error(err))

		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Status: utils.StatusResponse{
				Code:    http.StatusBadRequest,
				Message: "invalid request body",
			},
			Error: err.Error(),
		})
		return
	}

	category, err := h.categoryUseCase.FetchByID(id)
	if err != nil {
		h.logger.Error("fetchByID", zap.String("status", "error"), zap.String("method", "GET"), zap.Error(err))

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status: utils.StatusResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error while fetching category",
			},
			Error: err.Error(),
		})
		return
	}

	h.logger.Info("fetchByID", zap.String("status", "mapping"), zap.String("method", "GET"), zap.Any("category", category))

	result := dto.CategoryResponse{
		UUID:        category.UUID,
		Name:        category.Name,
		Description: category.Description,
	}

	h.logger.Info("fetchByID", zap.String("status", "fetched"), zap.String("method", "GET"), zap.Any("category", result))

	ctx.JSON(http.StatusOK, utils.SuccessResponse{
		Status: utils.StatusResponse{
			Code:    http.StatusOK,
			Message: "Category fetched",
		},
		Data: result,
	})
}
