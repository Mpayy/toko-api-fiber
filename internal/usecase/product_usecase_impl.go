package usecase

import (
	"context"
	"errors"
	"fmt"
	"toko-api-fiber/internal/entity"
	"toko-api-fiber/internal/exception"
	"toko-api-fiber/internal/model"
	"toko-api-fiber/internal/model/converter"
	"toko-api-fiber/internal/repository"

	"github.com/sirupsen/logrus"
)

type ProductUseCaseImpl struct {
	Transaction       repository.Transaction
	Log               *logrus.Logger
	ProductRepository repository.ProductRepository
}

func NewProductUseCase(tx repository.Transaction, log *logrus.Logger, productRepository repository.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{
		Transaction:       tx,
		Log:               log,
		ProductRepository: productRepository,
	}
}

func (u *ProductUseCaseImpl) Create(ctx context.Context, request *model.CreateProductRequest) (*model.ProductResponse, error) {
	product := &entity.Product{
		Name:  request.Name,
		Price: request.Price,
		Stock: *request.Stock,
	}

	err := u.ProductRepository.Create(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return converter.ToProductResponse(product), nil
}

func (u *ProductUseCaseImpl) Update(ctx context.Context, request *model.UpdateProductRequest) (*model.ProductResponse, error) {
	product, err := u.ProductRepository.GetByID(ctx, request.ID)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, fmt.Errorf("failed to find product by id: %w", err)
	}

	product.Name = request.Name
	product.Price = request.Price
	product.Stock = *request.Stock
	err = u.ProductRepository.Update(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return converter.ToProductResponse(product), nil
}

func (u *ProductUseCaseImpl) Delete(ctx context.Context, id int64) error {
	product, err := u.ProductRepository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return exception.ErrNotFound
		}
		return fmt.Errorf("failed to find product by id: %w", err)
	}

	err = u.ProductRepository.Delete(ctx, product)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}

func (u *ProductUseCaseImpl) GetAll(ctx context.Context) ([]*model.ProductResponse, error) {
	products, err := u.ProductRepository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all products: %w", err)
	}

	return converter.ToProductResponses(products), nil
}

func (u *ProductUseCaseImpl) GetByID(ctx context.Context, id int64) (*model.ProductResponse, error) {
	product, err := u.ProductRepository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, fmt.Errorf("failed to find product by id: %w", err)
	}

	return converter.ToProductResponse(product), nil
}

func (u *ProductUseCaseImpl) Patch(ctx context.Context, request *model.PatchProductRequest) (*model.ProductResponse, error) {
	product, err := u.ProductRepository.GetByID(ctx, request.ID)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, fmt.Errorf("failed to find product by id: %w", err)
	}

	fields := make(map[string]any)
	if request.Name != nil {
		fields["name"] = *request.Name
	}
	if request.Price != nil {
		fields["price"] = *request.Price
	}
	if request.Stock != nil {
		fields["stock"] = *request.Stock
	}

	if err := u.ProductRepository.Patch(ctx, product, fields); err != nil {
		return nil, fmt.Errorf("failed to patch product: %w", err)
	}

	return converter.ToProductResponse(product), nil
}
