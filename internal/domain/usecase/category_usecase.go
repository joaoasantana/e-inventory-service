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

	if _, err := uc.repository.FindByName(category.Name); err == nil {
		err = errors.New("category already exists")
		uc.logger.Info("create", zap.String("status", "creating"), zap.Any("category", category), zap.Error(err))

		return uuid.Nil, err
	}

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

		return nil, errors.New("categories not found")
	}

	if categories == nil || len(categories) == 0 {
		err = errors.New("categories not found")
		uc.logger.Error("fetchAll", zap.String("status", "error"), zap.Error(err))

		return nil, err
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

func (uc *CategoryUseCase) FetchByID(id string) (*model.Category, error) {
	uc.logger.Info("FetchByID", zap.String("status", "fetching"), zap.String("id", id))

	parsedID, err := uuid.Parse(id)
	if err != nil {
		uc.logger.Error("FetchByID", zap.String("status", "error"), zap.String("id", id), zap.Error(err))

		return nil, errors.New("failed to parse uuid")
	}

	category, err := uc.repository.FindByID(parsedID)
	if err != nil {
		uc.logger.Error("FetchByID", zap.String("status", "error"), zap.String("id", id), zap.Error(err))

		return nil, errors.New("category not found")
	}

	if category == nil {
		err = errors.New("category not found")
		uc.logger.Error("FetchByID", zap.String("status", "error"), zap.String("id", id), zap.Error(err))

		return nil, err
	}

	uc.logger.Info("FetchByID", zap.String("status", "mapping"), zap.Any("category", category))

	result := model.Category{
		UUID:        category.UUID,
		Name:        category.Name,
		Description: category.Description,
	}

	uc.logger.Info("FetchByID", zap.String("status", "fetched"), zap.Any("category", result))

	return &result, nil
}
