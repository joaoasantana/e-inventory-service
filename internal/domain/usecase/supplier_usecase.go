package usecase

import (
	"errors"
	"github.com/google/uuid"
	"github.com/joaoasantana/e-inventory-service/internal/domain/entity"
	"github.com/joaoasantana/e-inventory-service/internal/domain/model"
	"github.com/joaoasantana/e-inventory-service/internal/domain/repository"
	"go.uber.org/zap"
)

var (
	errGenerateUUID = errors.New("failed to generate uuid")
	errParseUUID    = errors.New("failed to parse uuid")

	errSupplierCreate    = errors.New("failed to create supplier")
	errSupplierFetchAll  = errors.New("failed to fetch all suppliers")
	errSupplierFetchByID = errors.New("failed to find supplier by uuid")

	errSupplierAlreadyExists = errors.New("supplier already exists")
	errSupplierNotFound      = errors.New("supplier not found")
	errSuppliersNotFound     = errors.New("suppliers not found")
)

type SupplierUseCase struct {
	Logger     *zap.Logger
	Repository repository.SupplierRepository
}

func (uc *SupplierUseCase) Create(supplier *model.Supplier) (uuid.UUID, error) {
	mLogger := uc.Logger.With(zap.String("method", "create"))

	mLogger.Info("creating supplier", zap.Any("supplier", supplier))

	if _, err := uc.Repository.FindByName(supplier.Name); err == nil {
		mLogger.Error("supplier already exists", zap.Error(err))
		return uuid.Nil, errSupplierAlreadyExists
	}

	id, err := uuid.NewUUID()
	if err != nil {
		mLogger.Error("failed to generate uuid", zap.Error(err))
		return uuid.Nil, errGenerateUUID
	}

	supplierEntity := &entity.Supplier{
		UUID:    id,
		Name:    supplier.Name,
		Contact: supplier.Contact,
	}

	if err = supplierEntity.ValidateRules(); err != nil {
		mLogger.Error("failed to validate rules", zap.Error(err))
		return uuid.Nil, err
	}

	if err = uc.Repository.Create(supplierEntity); err != nil {
		mLogger.Error("failed to create supplier", zap.Error(err))
		return uuid.Nil, errSupplierCreate
	}

	mLogger.Info("supplier created", zap.Any("supplier", supplierEntity))
	return supplierEntity.UUID, nil
}

func (uc *SupplierUseCase) FetchAll() ([]model.Supplier, error) {
	mLogger := uc.Logger.With(zap.String("method", "fetchAll"))

	mLogger.Info("fetching all suppliers")

	suppliers, err := uc.Repository.FindAll()
	if err != nil {
		mLogger.Error("failed to fetch all suppliers", zap.Error(err))
		return nil, errSupplierFetchAll
	}

	if suppliers == nil || len(suppliers) == 0 {
		mLogger.Error("suppliers is empty", zap.Error(errSuppliersNotFound))
		return nil, errSuppliersNotFound
	}

	var result []model.Supplier

	for _, supplier := range suppliers {
		result = append(result, model.Supplier{
			UUID:    supplier.UUID,
			Name:    supplier.Name,
			Contact: supplier.Contact,
		})
	}

	mLogger.Info("fetched all suppliers", zap.Any("suppliers", result))
	return result, nil
}

func (uc *SupplierUseCase) FetchByID(id string) (*model.Supplier, error) {
	mLogger := uc.Logger.With(zap.String("method", "fetchByID"))

	mLogger.Info("fetching supplier", zap.Any("id", id))

	parsedID, err := uuid.Parse(id)
	if err != nil {
		mLogger.Error("failed to parse supplier", zap.Error(errParseUUID))
		return nil, errParseUUID
	}

	supplier, err := uc.Repository.FindByID(parsedID)
	if err != nil {
		mLogger.Error("failed to fetch supplier", zap.Error(err))
		return nil, errSupplierFetchByID
	}

	if supplier == nil {
		mLogger.Error("supplier not found", zap.Error(errSupplierNotFound))
		return nil, errSupplierNotFound
	}

	result := &model.Supplier{
		UUID:    supplier.UUID,
		Name:    supplier.Name,
		Contact: supplier.Contact,
	}

	mLogger.Info("supplier found", zap.Any("supplier", result))
	return result, nil
}
