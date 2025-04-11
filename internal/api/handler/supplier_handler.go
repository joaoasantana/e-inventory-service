package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoasantana/e-inventory-service/internal/domain/usecase"
	"go.uber.org/zap"
)

type SupplierHandler struct {
	Logger  *zap.Logger
	UseCase *usecase.SupplierUseCase
}

func (h *SupplierHandler) CreateSupplier(ctx *gin.Context) {}

func (h *SupplierHandler) FetchAllSuppliers(ctx *gin.Context) {}

func (h *SupplierHandler) FetchSupplierByID(ctx *gin.Context) {}
