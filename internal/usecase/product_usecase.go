package usecase

import (
	"context"
	"toko-api-fiber/internal/model"
)

type ProductUsecase interface {
	Create(ctx context.Context, request *model.CreateProductRequest) (*model.ProductResponse, error)
	Update(ctx context.Context, request *model.UpdateProductRequest) (*model.ProductResponse, error)
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context, request *model.PaginationRequest) ([]*model.ProductResponse, int64, int, error)
	GetByID(ctx context.Context, id int64) (*model.ProductResponse, error)
	Patch(ctx context.Context, request *model.PatchProductRequest) (*model.ProductResponse, error)
}
