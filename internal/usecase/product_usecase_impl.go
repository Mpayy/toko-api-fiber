package usecase

import (
	"context"
	"errors"
	"fmt"
	"toko-api-fiber/internal/entity"
	"toko-api-fiber/internal/model"
	"toko-api-fiber/internal/model/converter"
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
		return nil, fmt.Errorf("%w: %s", model.ErrValidation, err.Error())
	}

	product := &entity.Product{
		Name:  request.Name,
		Price: request.Price,
		Stock: request.Stock,
	}

	if err := u.ProductRepository.Create(ctx, tx, product); err != nil {
		u.Log.WithError(err).Error("Failed to create product")
		return nil, fmt.Errorf("%w: %s", model.ErrInternal, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("Failed to commit transaction")
		return nil, fmt.Errorf("%w: %s", model.ErrInternal, err.Error())
	}

	return converter.ToProductResponse(product), nil
}

func (u *ProductUseCaseImpl) Update(ctx context.Context, request *model.UpdateProductRequest) (*model.ProductResponse, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		u.Log.WithError(err).Error("Validation Error")
		return nil, fmt.Errorf("%w: %s", model.ErrValidation, err.Error())
	}

	product, err := u.ProductRepository.GetByID(ctx, u.DB.WithContext(ctx), request.ID)
	if err != nil {
		u.Log.WithError(err).Error("Product not found")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: Product with id = %d not found", model.ErrNotFound, request.ID)
		}
		return nil, fmt.Errorf("%w: %s", model.ErrInternal, err.Error())
	}

	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	product.Name = request.Name
	product.Price = request.Price
	product.Stock = request.Stock

	if err := u.ProductRepository.Update(ctx, tx, product); err != nil {
		u.Log.WithError(err).Error("Failed to update product")
		return nil, fmt.Errorf("%w: %s", model.ErrInternal, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("Failed to commit transaction")
		return nil, fmt.Errorf("%w: %s", model.ErrInternal, err.Error())
	}

	return converter.ToProductResponse(product), nil
}

func (u *ProductUseCaseImpl) Delete(ctx context.Context, id int64) error {
	product, err := u.ProductRepository.GetByID(ctx, u.DB.WithContext(ctx), id)
	if err != nil {
		u.Log.WithError(err).Error("Product not found")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("%w: Product with id = %d not found", model.ErrNotFound, id)
		}
		return fmt.Errorf("%w: %s", model.ErrInternal, err.Error())
	}

	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.ProductRepository.Delete(ctx, tx, product); err != nil {
		u.Log.WithError(err).Error("Failed to delete product")
		return fmt.Errorf("%w: %s", model.ErrInternal, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("Failed to commit transaction")
		return fmt.Errorf("%w: %s", model.ErrInternal, err.Error())
	}

	return nil
}

func (u *ProductUseCaseImpl) GetAll(ctx context.Context) ([]*model.ProductResponse, error) {
	products, err := u.ProductRepository.GetAll(ctx, u.DB.WithContext(ctx))
	if err != nil {
		u.Log.WithError(err).Error("Failed to get all products")
		return nil, fmt.Errorf("%w: %s", model.ErrInternal, err.Error())
	}

	return converter.ToProductResponses(products), nil
}

func (u *ProductUseCaseImpl) GetByID(ctx context.Context, id int64) (*model.ProductResponse, error) {
	product, err := u.ProductRepository.GetByID(ctx, u.DB.WithContext(ctx), id)
	if err != nil {
		u.Log.WithError(err).Error("Product not found")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: Product with id = %d not found", model.ErrNotFound, id)
		}
		return nil, fmt.Errorf("%w: %s", model.ErrInternal, err.Error())
	}

	return converter.ToProductResponse(product), nil
}

func (u *ProductUseCaseImpl) Patch(ctx context.Context, request *model.PatchProductRequest) (*model.ProductResponse, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		u.Log.WithError(err).Error("Validation Error")
		return nil, fmt.Errorf("%w: %s", model.ErrValidation, err.Error())
	}

	product, err := u.ProductRepository.GetByID(ctx, u.DB.WithContext(ctx), request.ID)
	if err != nil {
		u.Log.WithError(err).Error("Product not found")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: Product with id = %d not found", model.ErrNotFound, request.ID)
		}
		return nil, fmt.Errorf("%w: %s", model.ErrInternal, err.Error())
	}

	fields := make(map[string]any)
	if request.Name != "" {
		fields["name"] = request.Name
	}
	if request.Price > 0 {
		fields["price"] = request.Price
	}
	if request.Stock >= 0 {
		fields["stock"] = request.Stock
	}

	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.ProductRepository.Patch(ctx, tx, product, fields); err != nil {
		u.Log.WithError(err).Error("Failed to patch product")
		return nil, fmt.Errorf("%w: %s", model.ErrInternal, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("Failed to commit transaction")
		return nil, fmt.Errorf("%w: %s", model.ErrInternal, err.Error())
	}

	return converter.ToProductResponse(product), nil
}
