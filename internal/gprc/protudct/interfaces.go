package protudct

import (
	"context"
	"server/internal/domein/models"
)
type Product interface {
	CreateUser(ctx context.Context, req *models.CreateProductReq)(*models.Product,error)
	DeleteProduct(ctx context.Context, req *models.ProductDeleteReq)(*models.Product ,error)
	GetAllProduct(ctx context.Context , req *models.GetAllProductReq) ([]*models.Product,error)
	GetProductById(ctx context.Context, id string)(*models.Product,error)
	UpdateProduc(ctx context.Context, req *models.UpdateProducReq)(*models.Product,error)
}