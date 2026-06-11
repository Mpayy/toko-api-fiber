package converter

import (
	"toko-api-fiber/internal/entity"
	"toko-api-fiber/internal/model"
)

func ToProductResponse(product *entity.Product) *model.ProductResponse {
	return &model.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func ToProductResponses(products []*entity.Product) []*model.ProductResponse {
	var productResponses []*model.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
