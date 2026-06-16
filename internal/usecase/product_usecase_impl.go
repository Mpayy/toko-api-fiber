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

type ProductUsecaseImpl struct {
	Transaction       repository.Transaction
	Log               *logrus.Logger
	ProductRepository repository.ProductRepository
}

func NewProductUseCase(tx repository.Transaction, log *logrus.Logger, productRepository repository.ProductRepository) ProductUsecase {
	return &ProductUsecaseImpl{
		Transaction:       tx,
		Log:               log,
		ProductRepository: productRepository,
	}
}

func (u *ProductUsecaseImpl) Create(ctx context.Context, request *model.CreateProductRequest) (*model.ProductResponse, error) {
	var stock int64
	if request.Stock != nil {
		stock = *request.Stock
	}
	product := &entity.Product{
		Name:  request.Name,
		Price: request.Price,
		Stock: stock,
	}

	if err := u.ProductRepository.Create(ctx, product); err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return converter.ToProductResponse(product), nil
}

func (u *ProductUsecaseImpl) Update(ctx context.Context, request *model.UpdateProductRequest) (*model.ProductResponse, error) {
	product, err := u.ProductRepository.GetByID(ctx, request.ID)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, fmt.Errorf("failed to find product by id: %w", err)
	}

	if request.Name != "" {
		product.Name = request.Name
	}
	if request.Price > 0 {
		product.Price = request.Price
	}
	if request.Stock != nil {
		product.Stock = *request.Stock
	}

	if err := u.ProductRepository.Update(ctx, product); err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return converter.ToProductResponse(product), nil
}

func (u *ProductUsecaseImpl) Delete(ctx context.Context, id int64) error {
	product, err := u.ProductRepository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return exception.ErrNotFound
		}
		return fmt.Errorf("failed to find product by id: %w", err)
	}

	if err := u.ProductRepository.Delete(ctx, product); err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}

func (u *ProductUsecaseImpl) GetAll(ctx context.Context, request *model.PaginationRequest) ([]*model.ProductResponse, int64, int, error) {
	page := 1
	if request.Page != nil && *request.Page > 0 {
		page = *request.Page
	}

	size := 10
	if request.Size != nil && *request.Size > 0 {
		size = *request.Size
	}

	products, totalItems, totalPage, err := u.ProductRepository.GetAll(ctx, page, size)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to get all products: %w", err)
	}

	productResponses := converter.ToProductResponses(products)

	return productResponses, totalItems, totalPage, nil
}

func (u *ProductUsecaseImpl) GetByID(ctx context.Context, id int64) (*model.ProductResponse, error) {
	product, err := u.ProductRepository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, fmt.Errorf("failed to find product by id: %w", err)
	}

	return converter.ToProductResponse(product), nil
}

func (u *ProductUsecaseImpl) Patch(ctx context.Context, request *model.PatchProductRequest) (*model.ProductResponse, error) {
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
