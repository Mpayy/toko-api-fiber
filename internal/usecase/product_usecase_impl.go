package usecase

import (
	"context"
	"errors"
	"fmt"
	"toko-api-fiber/internal/entity"
	"toko-api-fiber/internal/model"
	"toko-api-fiber/internal/model/converter"
	"toko-api-fiber/internal/repository"
	"toko-api-fiber/internal/util"

	"github.com/sirupsen/logrus"
)

type ProductUseCaseImpl struct {
	Transaction       util.Transaction
	Log               *logrus.Logger
	ProductRepository repository.ProductRepository
}

func NewProductUseCase(tx util.Transaction, log *logrus.Logger, productRepository repository.ProductRepository) ProductUseCase {
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
		Stock: request.Stock,
	}

	err := u.Transaction.WithTransaction(ctx, func(txCtx context.Context) error {
		if err := u.ProductRepository.Create(txCtx, product); err != nil {
			u.Log.WithError(err).Error("Failed to create product")
			return model.ErrInternal
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return converter.ToProductResponse(product), nil
}

func (u *ProductUseCaseImpl) Update(ctx context.Context, request *model.UpdateProductRequest) (*model.ProductResponse, error) {
	product, err := u.ProductRepository.GetByID(ctx, request.ID)
	if err != nil {
		u.Log.WithFields(logrus.Fields{
			"product_id": request.ID,
			"error":      err,
		}).Warn("Product not found")
		if errors.Is(err, model.ErrNotFound) {
			return nil, &model.NotFoundError{
				Message:    fmt.Sprintf("Product with id = %d not found", request.ID),
			}
		}
		return nil, model.ErrInternal
	}

	err = u.Transaction.WithTransaction(ctx, func(txCtx context.Context) error {
		product.Name = request.Name
		product.Price = request.Price
		product.Stock = request.Stock
		if err := u.ProductRepository.Update(txCtx, product); err != nil {
			u.Log.WithError(err).Error("Failed to update product")
			return model.ErrInternal
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return converter.ToProductResponse(product), nil
}

func (u *ProductUseCaseImpl) Delete(ctx context.Context, id int64) error {
	product, err := u.ProductRepository.GetByID(ctx, id)
	if err != nil {
		u.Log.WithFields(logrus.Fields{
			"product_id": id,
			"error":      err,
		}).Warn("Product not found")
		if errors.Is(err, model.ErrNotFound) {
			return &model.NotFoundError{
				Message:    fmt.Sprintf("Product with id = %d not found", id),
			}
		}
		return model.ErrInternal
	}

	err = u.Transaction.WithTransaction(ctx, func(txCtx context.Context) error {
		if err := u.ProductRepository.Delete(txCtx, product); err != nil {
			u.Log.WithError(err).Error("Failed to delete product")
			return model.ErrInternal
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (u *ProductUseCaseImpl) GetAll(ctx context.Context) ([]*model.ProductResponse, error) {
	products, err := u.ProductRepository.GetAll(ctx)
	if err != nil {
		u.Log.WithError(err).Error("Failed to get all products")
		return nil, model.ErrInternal
	}

	return converter.ToProductResponses(products), nil
}

func (u *ProductUseCaseImpl) GetByID(ctx context.Context, id int64) (*model.ProductResponse, error) {
	product, err := u.ProductRepository.GetByID(ctx, id)
	if err != nil {
		u.Log.WithFields(logrus.Fields{
			"product_id": id,
			"error":      err,
		}).Warn("Product not found")
		if errors.Is(err, model.ErrNotFound) {
			return nil, &model.NotFoundError{
				Message:    fmt.Sprintf("Product with id = %d not found", id),
			}
		}
		return nil, model.ErrInternal
	}

	return converter.ToProductResponse(product), nil
}

func (u *ProductUseCaseImpl) Patch(ctx context.Context, request *model.PatchProductRequest) (*model.ProductResponse, error) {
	product, err := u.ProductRepository.GetByID(ctx, request.ID)
	if err != nil {
		u.Log.WithFields(logrus.Fields{
			"product_id": request.ID,
			"error":      err,
		}).Warn("Product not found")
		if errors.Is(err, model.ErrNotFound) {
			return nil, &model.NotFoundError{
				Message:    fmt.Sprintf("Product with id = %d not found", request.ID),
			}
		}
		return nil, model.ErrInternal
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

	err = u.Transaction.WithTransaction(ctx, func(txCtx context.Context) error {
		if err := u.ProductRepository.Patch(txCtx, product, fields); err != nil {
			u.Log.WithError(err).Error("Failed to patch product")
			return model.ErrInternal
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return converter.ToProductResponse(product), nil
}
