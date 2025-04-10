package usecase

import (
	"errors"
	"github.com/google/uuid"
	"github.com/joaoasantana/e-inventory-service/internal/domain/entity"
	"github.com/joaoasantana/e-inventory-service/internal/domain/model"
	"github.com/joaoasantana/e-inventory-service/internal/domain/repository"
	"go.uber.org/zap"
)

type ProductUseCase struct {
	logger  *zap.Logger
	product repository.ProductRepository
}

func NewProductUseCase(logger *zap.Logger, product repository.ProductRepository) *ProductUseCase {
	mLogger := logger.With(
		zap.String("type", "service"),
		zap.String("domain", "product"),
	)

	return &ProductUseCase{mLogger, product}
}

func (uc *ProductUseCase) Create(product *model.Product) (uuid.UUID, error) {
	mLogger := uc.logger.With(zap.String("method", "create"))

	if _, err := uc.product.FindByName(product.Name); err == nil {
		mLogger.Error("error", zap.Error(errors.New("product already exists"))) // todo
		return uuid.Nil, errors.New("product already exists")                   // todo
	}

	id, err := uuid.NewUUID()
	if err != nil {
		mLogger.Error("error", zap.Error(err))
		return uuid.Nil, errors.New("failed to generate uuid") // todo
	}

	productEntity := &entity.Product{
		UUID:        id,
		Name:        product.Name,
		Image:       product.Image,
		Price:       product.Price,
		Description: product.Description,
	}

	if err = productEntity.ValidateRules(); err != nil {
		mLogger.Error("error", zap.Error(err))
		return uuid.Nil, err
	}

	if err = uc.product.Create(productEntity); err != nil {
		mLogger.Error("error", zap.Error(err))
		return uuid.Nil, errors.New("failed to create product")
	}

	mLogger.Info("success", zap.Error(err), zap.String("id", id.String()))
	return id, nil
}

func (uc *ProductUseCase) FetchAll() ([]model.Product, error) {
	mLogger := uc.logger.With(zap.String("method", "fetchAll"))

	products, err := uc.product.FindAll()
	if err != nil {
		mLogger.Error("error", zap.Error(err))
		return nil, err
	}

	if products == nil || len(products) == 0 {
		mLogger.Error("error", zap.Error(errors.New("products is empty"))) // todo
		return nil, errors.New("products is empty")                        // todo
	}

	var result []model.Product

	for _, product := range products {
		result = append(result, model.Product{
			UUID:        product.UUID,
			Name:        product.Name,
			Image:       product.Image,
			Price:       product.Price,
			Description: product.Description,
		})
	}

	mLogger.Info("success", zap.Any("products", result))
	return result, nil
}

func (uc *ProductUseCase) FetchByID(id string) (*model.Product, error) {
	mLogger := uc.logger.With(zap.String("method", "fetchByID"))

	parsedID, err := uuid.Parse(id)
	if err != nil {
		mLogger.Error("error", zap.String("id", id), zap.Error(err))
		return nil, errors.New("failed to parse uuid") // todo
	}

	product, err := uc.product.FindByID(parsedID)
	if err != nil {
		mLogger.Error("error", zap.String("id", id), zap.Error(err))
		return nil, err
	}

	if product == nil {
		mLogger.Error("error", zap.String("id", id), zap.Error(errors.New("product not found"))) // todo
		return nil, errors.New("product not found")                                              // todo
	}

	result := &model.Product{
		UUID:        product.UUID,
		Name:        product.Name,
		Image:       product.Image,
		Price:       product.Price,
		Description: product.Description,
	}

	mLogger.Info("success", zap.Any("product", result))
	return result, nil
}
