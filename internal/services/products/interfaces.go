package products

import (
	"context"
	"server/internal/domein/models"
)

type SaverProduct interface {
	SaveProduct(ctx context.Context,req *models.CreateProductReq)(*models.Product,error)
}
type ProviderProduct interface {
	GetProdcutById(ctx context.Context, id string)(*models.Product,error)
	GetAllProduct(ctx context.Context, 
		req *models.GetAllProductReq) ([]*models.Product,error)
}

type UpdaterProduct interface {
	UpdateProduct(ctx context.Context, 
		req *models.UpdateProducReq) (*models.Product,error)
}
type DeleterProduct interface {
	DeleteProduct(ctx context.Context,  req *models.ProductDeleteReq)(*models.Product,error)
}