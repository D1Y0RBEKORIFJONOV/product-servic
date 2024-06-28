package postgresql

import (
	"context"
	"fmt"
	"server/internal/domein/models"
	"server/internal/pkg/postgres"
)

// type SaverProduct interface {
// 	SaveProduct(ctx context.Context,req *models.CreateProductReq)(*models.Product,error)
// }
// type ProviderProduct interface {
// 	GetProdcutById(ctx context.Context, id string)(*models.Product,error)
// 	GetAllProduct(ctx context.Context,
// 		req *models.GetAllProductReq) ([]*models.Product,error)
// }

// type UpdaterProduct interface {
// 	UpdateProduct(ctx context.Context,
// 		req *models.UpdateProducReq) (*models.Product,error)
// }
// type DeleterProduct interface {
// 	DeleteProduct(ctx context.Context,  req *models.ProductDeleteReq)(*models.Product,error)
// }

type ProductRepository struct {
	db        *postgres.PostgresDB
	tableName string
}

func NewProductRepository(pg *postgres.PostgresDB) *ProductRepository {
	return &ProductRepository{
		db:        pg,
		tableName: "products",
	}
}

func (p *ProductRepository)selectQueryProduct()string{
	return `
	id,
	name_product,
	category,
	price,
	count,
	status_product,
	created_at,
	updated_at,
	deleted_at
	`
}

func (p *ProductRepository) SaveProduct(ctx context.Context, req *models.CreateProductReq) (*models.Product, error) {
	data := map[string]interface{}{
		"name_product":req.Name,
		"category":req.Category,
		"price":req.Price,
		"count":req.Count,
		"status":"created_at",
	}
	query,argc,err := p.db.Sq.Builder.Insert(p.tableName).
	SetMap(data).
	Suffix(fmt.Sprintf("RETURNING %s",p.selectQueryProduct())).ToSql()

	if err != nil {
		return nil,err
	}
	var product models.Product
	err  = p.db.QueryRow(ctx,query,argc...).Scan(
		&product.Id,
		&product.Name,
		&product.Category,
		&product.Count,
		&product.Status,
		&product.Created_at,
		&product.Updated_at,
		&product.Deleted_at,
	)
	if err != nil {
		return nil,err
	}
	return &product,nil
}
