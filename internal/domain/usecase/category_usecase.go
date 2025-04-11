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
	Logger     *zap.Logger
	Repository repository.CategoryRepository
}

func (uc *CategoryUseCase) Create(category *model.Category) (uuid.UUID, error) {
	mLogger := uc.Logger.With(zap.String("method", "create"))

	if _, err := uc.Repository.FindByName(category.Name); err == nil {
		mLogger.Error("error", zap.Error(errors.New("category already exists"))) // todo
		return uuid.Nil, errors.New("category already exists")                   // todo
	}

	id, err := uuid.NewUUID()
	if err != nil {
		mLogger.Error("error", zap.Error(err))
		return uuid.Nil, errors.New("failed to generate uuid") // todo
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

	if err = uc.Repository.Create(categoryEntity); err != nil {
		mLogger.Error("error", zap.Error(err))
		return uuid.Nil, errors.New("failed to create category")
	}

	mLogger.Info("success", zap.Error(err), zap.String("id", id.String()))
	return id, nil
}

func (uc *CategoryUseCase) FetchAll() ([]model.Category, error) {
	mLogger := uc.Logger.With(zap.String("method", "fetchAll"))

	categories, err := uc.Repository.FindAll()
	if err != nil {
		mLogger.Error("error", zap.Error(err))
		return nil, err
	}

	if categories == nil || len(categories) == 0 {
		mLogger.Error("error", zap.Error(errors.New("categories is empty"))) // todo
		return nil, errors.New("categories is empty")                        // todo
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
	mLogger := uc.Logger.With(zap.String("method", "findByID"))

	parsedID, err := uuid.Parse(id)
	if err != nil {
		mLogger.Error("error", zap.String("id", id), zap.Error(err))
		return nil, errors.New("failed to parse uuid") // todo
	}

	category, err := uc.Repository.FindByID(parsedID)
	if err != nil {
		mLogger.Error("error", zap.String("id", id), zap.Error(err))
		return nil, err
	}

	if category == nil {
		mLogger.Error("error", zap.String("id", id), zap.Error(errors.New("category not found"))) // todo
		return nil, errors.New("category not found")                                              // todo
	}

	result := model.Category{
		UUID:        category.UUID,
		Name:        category.Name,
		Description: category.Description,
	}

	mLogger.Info("success", zap.Any("category", result))
	return &result, nil
}
