package usecase

import (
	"errors"
	"github.com/google/uuid"
	"github.com/joaoasantana/e-inventory-service/internal/domain/entity"
	"github.com/joaoasantana/e-inventory-service/internal/domain/model"
	"github.com/joaoasantana/e-inventory-service/internal/domain/repository"
	"go.uber.org/zap"
)

type CategoryUseCase struct {
	logger     *zap.Logger
	repository repository.CategoryRepository
}

func NewCategoryUseCase(logger *zap.Logger, repository repository.CategoryRepository) *CategoryUseCase {
	mLogger := logger.With(
		zap.String("type", "service"),
		zap.String("domain", "category"),
	)

	return &CategoryUseCase{mLogger, repository}
}

func (uc *CategoryUseCase) Create(category *model.Category) (uuid.UUID, error) {
	mLogger := uc.logger.With(zap.String("method", "create"))

	if _, err := uc.repository.FindByName(category.Name); err == nil {
		mLogger.Error("error", zap.Error(errors.New("category already exists")))
		return uuid.Nil, errors.New("category already exists")
	}

	id, err := uuid.NewUUID()
	if err != nil {
		mLogger.Error("error", zap.Error(err))
		return uuid.Nil, errors.New("failed to generate uuid")
	}

	categoryEntity := &entity.Category{
		UUID:        id,
		Name:        category.Name,
		Description: category.Description,
	}

	if err = categoryEntity.ValidateRules(); err != nil {
		mLogger.Error("error", zap.Error(err))
		return uuid.Nil, err
	}

	if err = uc.repository.Create(categoryEntity); err != nil {
		mLogger.Error("error", zap.Error(err))
		return uuid.Nil, errors.New("failed to create category")
	}

	mLogger.Info("success", zap.Error(err), zap.String("id", id.String()))
	return id, nil
}

func (uc *CategoryUseCase) FetchAll() ([]model.Category, error) {
	mLogger := uc.logger.With(zap.String("method", "fetchAll"))

	categories, err := uc.repository.FindAll()
	if err != nil {
		mLogger.Error("error", zap.Error(errors.New("categories not found")))
		return nil, errors.New("categories not found")
	}

	if categories == nil || len(categories) == 0 {
		mLogger.Error("error", zap.Error(errors.New("categories is empty")))
		return nil, errors.New("categories is empty")
	}

	var result []model.Category

	for _, category := range categories {
		result = append(result, model.Category{
			UUID:        category.UUID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	mLogger.Info("success", zap.Any("categories", result))
	return result, nil
}

func (uc *CategoryUseCase) FetchByID(id string) (*model.Category, error) {
	mLogger := uc.logger.With(zap.String("method", "findByID"))

	parsedID, err := uuid.Parse(id)
	if err != nil {
		mLogger.Error("error", zap.String("id", id), zap.Error(err))
		return nil, errors.New("failed to parse uuid")
	}

	category, err := uc.repository.FindByID(parsedID)
	if err != nil {
		mLogger.Error("error", zap.String("id", id), zap.Error(err))
		return nil, errors.New("category not found")
	}

	result := model.Category{
		UUID:        category.UUID,
		Name:        category.Name,
		Description: category.Description,
	}

	mLogger.Info("success", zap.Any("category", result))
	return &result, nil
}
