package usecase

import (
	"context"
	"errors"
	"toko-api-fiber/internal/entity"
	"toko-api-fiber/internal/model"
	"toko-api-fiber/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductUseCaseImpl struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	ProductRepository repository.ProductRepository
}

func NewProductUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, productRepository repository.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{
		DB:                db,
		Log:               log,
		Validate:          validate,
		ProductRepository: productRepository,
	}
}

func (u *ProductUseCaseImpl) Create(ctx context.Context, request *model.CreateProductRequest) (*model.ProductResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := u.Validate.Struct(request)
	if err != nil {
		u.Log.WithError(err).Error("Validation Error")
		return nil, errors.New("Validation Error")
	}

	product := &entity.Product{
		Name:  request.Name,
		Price: request.Price,
		Stock: request.Stock,
	}

	if err := u.ProductRepository.Create(ctx, tx, product); err != nil {
		u.Log.WithError(err).Error("Failed to create product")
		return nil, errors.New("Failed to create product")
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("Failed to commit transaction")
		return nil, errors.New("Failed to commit transaction")
	}

	response := &model.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	return response, nil
}

func (u *ProductUseCaseImpl) Update(ctx context.Context, request *model.UpdateProductRequest) (*model.ProductResponse, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		u.Log.WithError(err).Error("Validation Error")
		return nil, errors.New("Validation Error")
	}

	product, err := u.ProductRepository.GetByID(ctx, u.DB.WithContext(ctx), request.ID)
	if err != nil {
		u.Log.WithError(err).Error("Product not found")
		return nil, errors.New("Product not found")
	}

	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	product.Name = request.Name
	product.Price = request.Price
	product.Stock = request.Stock

	if err := u.ProductRepository.Update(ctx, tx, product); err != nil {
		u.Log.WithError(err).Error("Failed to update product")
		return nil, errors.New("Failed to update product")
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("Failed to commit transaction")
		return nil, errors.New("Failed to commit transaction")
	}

	response := &model.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	return response, nil
}

func (u *ProductUseCaseImpl) Delete(ctx context.Context, id int64) error {
	product, err := u.ProductRepository.GetByID(ctx, u.DB.WithContext(ctx), id)
	if err != nil {
		u.Log.WithError(err).Error("Product not found")
		return errors.New("Product not found")
	}

	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.ProductRepository.Delete(ctx, tx, product); err != nil {
		u.Log.WithError(err).Error("Failed to delete product")
		return errors.New("Failed to delete product")
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("Failed to commit transaction")
		return errors.New("Failed to commit transaction")
	}

	return nil
}

func (u *ProductUseCaseImpl) GetAll(ctx context.Context) ([]*model.ProductResponse, error) {
	products, err := u.ProductRepository.GetAll(ctx, u.DB.WithContext(ctx))
	if err != nil {
		u.Log.WithError(err).Error("Failed to get all products")
		return nil, errors.New("Failed to get all products")
	}

	response := make([]*model.ProductResponse, 0)
	for _, product := range products {
		response = append(response, &model.ProductResponse{
			ID:        product.ID,
			Name:      product.Name,
			Price:     product.Price,
			Stock:     product.Stock,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		})
	}

	return response, nil
}

func (u *ProductUseCaseImpl) GetByID(ctx context.Context, id int64) (*model.ProductResponse, error) {
	product, err := u.ProductRepository.GetByID(ctx, u.DB.WithContext(ctx), id)
	if err != nil {
		u.Log.WithError(err).Error("Product not found")
		return nil, errors.New("Product not found")
	}

	response := &model.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	return response, nil
}
