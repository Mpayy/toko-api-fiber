package usecase

import (
	"context"
	"toko-api-fiber/internal/model"
)

type ProductUseCase interface {
	Create(ctx context.Context, request *model.CreateProductRequest) (*model.ProductResponse, error)
	Update(ctx context.Context, request *model.UpdateProductRequest) (*model.ProductResponse, error)
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]*model.ProductResponse, error)
	GetByID(ctx context.Context, id int64) (*model.ProductResponse, error)
	Patch(ctx context.Context, request *model.PatchProductRequest) (*model.ProductResponse, error)
}