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
	uc.logger.Info("create", zap.String("status", "creating"), zap.Any("category", category))

	id, err := uuid.NewUUID()
	if err != nil {
		uc.logger.Error("create", zap.String("status", "error"), zap.Error(err))

		return uuid.Nil, errors.New("failed to generate uuid")
	}

	categoryEntity := &entity.Category{
		UUID:        id,
		Name:        category.Name,
		Description: category.Description,
	}

	if err = categoryEntity.ValidateRules(); err != nil {
		uc.logger.Error("create", zap.String("status", "error"), zap.Error(err))

		return uuid.Nil, err
	}

	if err = uc.repository.Create(categoryEntity); err != nil {
		uc.logger.Error("create", zap.String("status", "error"), zap.Error(err))

		return uuid.Nil, errors.New("failed to create category")
	}

	uc.logger.Info("create", zap.String("status", "created"), zap.Any("category", categoryEntity))

	return id, nil
}

func (uc *CategoryUseCase) FetchAll() ([]model.Category, error) {
	uc.logger.Info("fetchAll", zap.String("status", "fetching"))

	categories, err := uc.repository.FindAll()
	if err != nil {
		uc.logger.Error("fetchAll", zap.String("status", "error"), zap.Error(err))

		return nil, errors.New("failed to fetch all categories")
	}

	if categories == nil || len(categories) == 0 {
		uc.logger.Error("fetchAll", zap.String("status", "error"), zap.Error(errors.New("no categories found")))

		return nil, errors.New("categories is empty")
	}

	uc.logger.Info("fetchAll", zap.String("status", "mapping"), zap.Any("categories", categories))

	var result []model.Category

	for _, category := range categories {
		result = append(result, model.Category{
			UUID:        category.UUID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	uc.logger.Info("fetchAll", zap.String("status", "fetched"), zap.Any("categories", result))

	return result, nil
}
